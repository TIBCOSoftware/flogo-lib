package trigger

import (
	//"context"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/stretchr/testify/assert"
)

type MockFactory struct {
}

func (f *MockFactory) New(id string) Trigger2 {
	return &MockTrigger{}
}

type MockTrigger struct {
}

func (t *MockTrigger) Init(config types.TriggerConfig, actionRunner action.Runner) {
	//Noop
}

func (t *MockTrigger) Start() error { return nil }
func (t *MockTrigger) Stop() error  { return nil }

//TestAddFactoryEmptyRef
func TestAddFactoryEmptyRef(t *testing.T) {

	reg := &registry{}

	// Add factory
	err := reg.AddFactory("", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterFactory: ref is empty", err.Error())
}

//TestAddFactoryNilFactory
func TestAddFactoryNilFactory(t *testing.T) {

	reg := &registry{}

	// Add factory
	err := reg.AddFactory("github.com/mock", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterFactory: factory is nil", err.Error())
}

//TestAddFactoryDuplicated
func TestAddFactoryDuplicated(t *testing.T) {

	reg := &registry{}
	f := &MockFactory{}

	// Add factory: this time should pass
	err := reg.AddFactory("github.com/mock", f)
	assert.Nil(t, err)
	// Add factory: this time should fail, duplicated
	err = reg.AddFactory("github.com/mock", f)
	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterFactory: already registered factory for ref 'github.com/mock'", err.Error())
}

//TestAddFactoryOk
func TestAddFactoryOk(t *testing.T) {

	reg := &registry{}
	f := &MockFactory{}

	// Add factory
	err := reg.AddFactory("github.com/mock", f)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reg.factories))
}

//TestGetFactoriesOk
func TestGetFactoriesOk(t *testing.T) {

	reg := &registry{}
	f := &MockFactory{}

	// Add factory
	err := reg.AddFactory("github.com/mock", f)
	assert.Nil(t, err)

	// Get factory
	fs := reg.GetFactories()
	assert.Equal(t, 1, len(fs))
}

//TestAddInstanceEmptyId
func TestAddInstanceEmptyId(t *testing.T) {

	reg := &registry{}

	// Add factory
	err := reg.AddInstance("", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterInstance: id is empty", err.Error())
}

//TestAddInstanceNilInstance
func TestAddInstanceNilInstance(t *testing.T) {

	reg := &registry{}

	// Add instance
	err := reg.AddInstance("myInstanceId", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterInstance: instance is nil", err.Error())
}

//TestAddInstanceDuplicated
func TestAddInstanceDuplicated(t *testing.T) {

	reg := &registry{}
	i := &TriggerInstance{}

	// Add instance: this time should pass
	err := reg.AddInstance("myinstanceId", i)
	assert.Nil(t, err)
	// Add instance: this time should fail, duplicated
	err = reg.AddInstance("myinstanceId", i)
	assert.NotNil(t, err)
	assert.Equal(t, "registry.RegisterInstance: already registered instance for id 'myinstanceId'", err.Error())
}

//TestAddInstanceOk
func TestAddInstanceOk(t *testing.T) {

	reg := &registry{}
	i := &TriggerInstance{}

	// Add instance
	err := reg.AddInstance("myinstanceId", i)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reg.instances))
}
