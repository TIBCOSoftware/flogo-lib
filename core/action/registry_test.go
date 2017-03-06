package action

import (
	"context"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/stretchr/testify/assert"
)

type MockFactory struct {
}

func (f *MockFactory) New(id string) Action2 {
	return &MockAction{}
}

type MockAction struct {
}

func (a *MockAction) Run(context context.Context, uri string, options interface{}, handler ResultHandler) error {
	return nil
}

func (a *MockAction) Init(config types.ActionConfig) {
	//Noop
}

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
	i := &ActionInstance{}

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
	i := &ActionInstance{}

	// Add instance
	err := reg.AddInstance("myinstanceId", i)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reg.instances))
}

//TestGetActionOk
func TestGetActionOk(t *testing.T) {

	reg := &registry{}
	i := &ActionInstance{Interf: &MockAction{}}

	// Add instance
	err := reg.AddInstance("myinstanceId", i)
	assert.Nil(t, err)

	a := reg.GetAction("myinstanceId")
	assert.NotNil(t, a)

	a = reg.GetAction("myunknowninstanceId")
	assert.Nil(t, a)
}
