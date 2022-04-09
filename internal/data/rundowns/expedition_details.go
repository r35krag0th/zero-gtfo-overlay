package rundowns

type ExpeditionDetails struct {
	ID                        string          `yaml:"id"`
	Name                      string          `yaml:"name"`
	Depth                     int             `yaml:"depth"`
	Intel                     string          `yaml:"intel"`
	Extended                  bool            `yaml:"extended"`
	InterruptedCommunications string          `yaml:"interrupted_communications"`
	Sectors                   []SectorDetails `yaml:"sectors"`
}

func (e ExpeditionDetails) GetSector(name string) *SectorDetails {
	for _, sector := range e.Sectors {
		if sector.Name == name {
			return &sector
		}
	}
	return nil
}

func (e ExpeditionDetails) AvailableSectors() []string {
	output := make([]string, len(e.Sectors))
	for index, sector := range e.Sectors {
		output[index] = sector.Name
	}
	return output
}
