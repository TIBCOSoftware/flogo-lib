package activity

import "github.com/TIBCOSoftware/flogo-lib/flow/support"

// Context describes the execution context for an Activity.
// It provides access to attributes, task and Flow information.
type Context interface {

	// FlowDetails returns the details fo the Flow Instance
	FlowDetails() FlowDetails

	// TaskName returns the name of the Task the Activity is currently executing
	TaskName() string

	// GetInput gets the value of the specified input attribute
	GetInput(name string) interface{}

	// SetOutput sets the value of the specified output attribute
	SetOutput(name string, value interface{})
	
	// GetSetting returns the activity configuration value
	GetSetting(name string) interface{}
}

// FlowDetails details of the flow that is being executed
type FlowDetails interface {

	// ID returns the ID of the Flow Instance
	ID() string

	// FlowName returns the name of the Flow
	Name() string

	// ReplyHandler returns the reply handler for the flow Instance
	ReplyHandler() support.ReplyHandler
}
