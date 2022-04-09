package rundowns

import (
	"fmt"
	"io/ioutil"
	"os"

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

func LoadRundownDetailsFromYAML(filename string) (*RundownDetails, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var output RundownDetails
	err = yaml.Unmarshal(b, &output)
	return &output, err
}
