package app

import (
	"context"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/stretchr/testify/assert"
)

//TestCreateTriggersOk
func TestCreateTriggersOk(t *testing.T) {

	app := getMockApp()

	// Create the mock factories
	tFactories := make(map[string]trigger.Factory, 1)
	tFactories["github.com/TIBCOSoftware/flogo-lib/app/mocktrigger"] = &MockTriggerFactory{}

	helper := NewInstanceHelper(app, tFactories, nil)

	triggers, err := helper.CreateTriggers()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(triggers))
}

//TestCreateActionsOk
func TestCreateActionsOk(t *testing.T) {

	app := getMockApp()

	// Create the mock factories
	aFactories := make(map[string]action.Factory, 1)
	aFactories["github.com/TIBCOSoftware/flogo-lib/app/mockaction"] = &MockActionFactory{}

	helper := NewInstanceHelper(app, nil, aFactories)

	actions, err := helper.CreateActions()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(actions))
}

//MockTriggerFactory
type MockTriggerFactory struct {
}

//MockTrigger
type MockTrigger struct {
}

func (t *MockTrigger) Init(config types.TriggerConfig, actionRunner action.Runner) {
	//Noop
}
func (t *MockTrigger) Start() error {
	return nil
}
func (t *MockTrigger) Stop() error {
	return nil
}
func (t *MockTriggerFactory) New(id string) trigger.Trigger2 {
	return &MockTrigger{}
}

//MockActionFactory
type MockActionFactory struct {
}

//MockAction
type MockAction struct {
}

func (t *MockAction) Init(config types.ActionConfig) {
	//Noop
}
func (t *MockAction) Start() error {
	return nil
}
func (t *MockAction) Stop() error {
	return nil
}

func (t *MockAction) Run(context context.Context, uri string, options interface{}, handler action.ResultHandler) error {
	return nil
}

func (t *MockActionFactory) New(id string) action.Action2 {
	return &MockAction{}
}

//getMockApp returns a mock app
func getMockApp() *types.AppConfig {
	triggers := make([]*types.TriggerConfig, 1)

	trigger1 := &types.TriggerConfig{Id: "myTrigger1", Ref: "github.com/TIBCOSoftware/flogo-lib/app/mocktrigger"}
	triggers[0] = trigger1

	actions := make([]*types.ActionConfig, 1)

	action1 := &types.ActionConfig{Id: "myAction1", Ref: "github.com/TIBCOSoftware/flogo-lib/app/mockaction"}
	actions[0] = action1

	return &types.AppConfig{Name: "MyApp", Version: "1.0.0", Triggers: triggers, Actions: actions}
}
