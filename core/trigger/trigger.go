package trigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/TIBCOSoftware/flogo-lib/util"
)

// Trigger is object that triggers/starts flow instances and
// is managed by an engine
type Trigger interface {
	util.Managed

	// TriggerMetadata returns the metadata of the trigger
	Metadata() *Metadata

	// Init sets up the trigger, it is called before Start()
	Init(config *Config, actionRunner action.Runner)
}

// Factory is used to create new instances for a trigger
type Factory interface {
	New(id string) Trigger2
}

// TODO change name from Trigger2 to Trigger once app refactoring is done
// Trigger is object that triggers/starts flow instances and
// is managed by an engine
type Trigger2 interface {
	util.Managed

	// Init sets up the trigger, it is called before Start()
	Init(config types.TriggerConfig, actionRunner action.Runner)
}

//TriggerInstance contains all the information for a Trigger Instance, configuration and interface
type TriggerInstance struct {
	Config *types.TriggerConfig
	Interf Trigger2
}
