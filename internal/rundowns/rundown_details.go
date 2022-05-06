package rundowns

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
)

type RundownDetails struct {
	ID          int                 `yaml:"id"`
	Name        string              `yaml:"name"`
	Expeditions []ExpeditionDetails `yaml:"expeditions"`
}

func (r RundownDetails) ShortName() string {
	return fmt.Sprintf("R%d", r.ID)
}

func (r RundownDetails) LongName() string {
	return fmt.Sprintf("Rundown %d", r.ID)
}

func (r RundownDetails) GetExpedition(id string) *ExpeditionDetails {
	for _, expedition := range r.Expeditions {
		if expedition.ID == id {
			return &expedition
		}
	}
	return nil
}

func (r RundownDetails) HasExpedition(id string) bool {
	for _, expedition := range r.Expeditions {
		if expedition.ID == id {
			return true
		}
	}
	return false
}

func (r RundownDetails) IsExtendedExpedition(id string) (bool, error) {
	expedition := r.GetExpedition(id)
	if expedition == nil {
		return false, fmt.Errorf("unknown expedition for rundown %d: %s", r.ID, id)
	}
	return expedition.Extended, nil
}

func LoadFile(filename string) (*RundownDetails, error) {
	if _, err := os.Stat(filename); err != nil {
		logx.Errorf("rundown data file does not exist: %s", filename)
		return nil, err
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		logx.Errorf("failed to read rundown data file %s: %v", filename, err)
		return nil, err
	}

	var output RundownDetails
	err = yaml.Unmarshal(b, &output)
	return &output, err
}

func LoadFiles(filenames []string) (map[string]*RundownDetails, []error) {
	logx.Infof("attempting to load rundown data from %d files", len(filenames))
	output := make(map[string]*RundownDetails)
	var errorOutput []error
	for _, fileName := range filenames {
		details, err := LoadFile(fileName)
		if err != nil {
			errorOutput = append(errorOutput, fmt.Errorf("failed to load rundown data from %s: %v", fileName, err))
			logx.Errorf("failed to load rundown data from %s: %v", fileName, err)
			continue
		}

		output[strconv.Itoa(details.ID)] = details
		logx.Infof("loaded rundown data from %s as 'Rundown %d'", fileName, details.ID)
	}
	return output, errorOutput
}
