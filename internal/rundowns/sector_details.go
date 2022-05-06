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
	ActivateGeneratorClusterObjective SectorObjectiveKind = "Activate Generator Cluster"
	DistributePowercellsObjective     SectorObjectiveKind = "Distribute Powercells"
	EstablishUplinkObjective          SectorObjectiveKind = "Establish Uplink"
	GatherItemsObjective              SectorObjectiveKind = "Gather Items"
	ProcessItemObjective              SectorObjectiveKind = "Process Item"
	ReactorShutdownObjective          SectorObjectiveKind = "Reactor Shutdown"
	ReactorStartupObjective           SectorObjectiveKind = "Reactor Startup"
	RetrieveHSUObjective              SectorObjectiveKind = "Retrieve HSU"
	RetrieveItemObjective             SectorObjectiveKind = "Retrieve Item"
	SurviveWardenProtocolObjective    SectorObjectiveKind = "Survive Warden Protocol"
	TerminalCommandObjective          SectorObjectiveKind = "Terminal Command"
	UnknownObjective                  SectorObjectiveKind = "Unknown"
	UseItem                           SectorObjectiveKind = "Use Item"
)

type SectorDetails struct {
	Name       string                  `yaml:"name"`
	Objectives []SectorObjectiveDetail `yaml:"objectives"`
	// ObjectivesV2 []SectorObjectiveDetail `yaml:"objective_v2"`
}

type SectorObjectiveDetail struct {
	Kind SectorObjectiveKind `yaml:"kind"`
	// Kind     string `yaml:"kind"`
	Metadata []struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	} `yaml:"metadata"`
}
