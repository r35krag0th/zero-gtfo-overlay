package rundowns

type SectorName = string

const (
	SectorNameHigh               SectorName = "High"
	SectorNameExtreme            SectorName = "Extreme"
	SectorNameOverload           SectorName = "Overload"
	SectorNamePrisonerEfficiency SectorName = "Prisoner Efficiency"
	SectorNameMain               SectorName = "Main"
	SectorNameSecondary          SectorName = "Secondary"
)

type SectorObjectiveKind = string

const (
	TerminalCommandObjective          SectorObjectiveKind = "Terminal Command"
	GatherItemsObjective              SectorObjectiveKind = "Gather Items"
	RetrieveItemObjective             SectorObjectiveKind = "Retrieve Item"
	RetrieveHSUObjective              SectorObjectiveKind = "Retrieve HSU"
	DistributePowercellsObjective     SectorObjectiveKind = "Distribute Powercells"
	ReactorStartupObjective           SectorObjectiveKind = "Reactor Startup"
	ReactorShutdownObjective          SectorObjectiveKind = "Reactor Shutdown"
	EstablishUplinkObjective          SectorObjectiveKind = "Establish Uplink"
	ProcessItemObjective              SectorObjectiveKind = "Process Item"
	SurviveWardenProtocolObjective    SectorObjectiveKind = "Survive Warden Protocol"
	ActivateGeneratorClusterObjective SectorObjectiveKind = "Activate Generator Cluster"
)

type SectorDetails struct {
	Name       string                  `yaml:"name"`
	Objectives []SectorObjectiveDetail `yaml:"objectives"`
	// ObjectivesV2 []SectorObjectiveDetail `yaml:"objective_v2"`
}

type SectorObjectiveDetail struct {
	Kind     SectorObjectiveKind `yaml:"kind"`
	Metadata []struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	} `yaml:"metadata"`
}
