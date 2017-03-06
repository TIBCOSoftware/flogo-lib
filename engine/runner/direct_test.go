package runner

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
)

type MockAction struct {
	mock.Mock
}

func (m *MockAction) Run(context context.Context, uri string, options interface{}, handler action.ResultHandler) error {
	args := m.Called(context, uri, options)
	if handler != nil {
		handler.HandleResult(0, "mock", nil)
		handler.Done()
	}
	return args.Error(0)
}

//Test that Result returns the expected values
func TestResultOk(t *testing.T) {
	rh := &SyncResultHandler{code: 1, data: "mock data", err: errors.New("New Error")}
	code, data, err := rh.Result()
	assert.Equal(t, 1, code)
	assert.Equal(t, "mock data", data)
	assert.NotNil(t, err)
}

//Test Direct Start method
func TestDirectStartOk(t *testing.T) {
	runner := NewDirect()
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
}

//Test Stop method
func TestDirectStopOk(t *testing.T) {
	runner := NewDirect()
	assert.NotNil(t, runner)
	err := runner.Stop()
	assert.Nil(t, err)
}

//Test Run method with a nil action
func TestDirectRunNilAction(t *testing.T) {
	runner := NewDirect()
	assert.NotNil(t, runner)
	_, _, err := runner.Run(nil, nil, "", nil)
	assert.NotNil(t, err)
}

//Test Run method with error running action
func TestDirectRunErr(t *testing.T) {
	runner := NewDirect()
	assert.NotNil(t, runner)
	// Mock Action
	mockAction := new(MockAction)
	mockAction.On("Run", nil, "", nil).Return(errors.New("Action Error"))
	_, _, err := runner.Run(nil, mockAction, "", nil)
	assert.NotNil(t, err)
}

//Test Run method ok
func TestDirectRunOk(t *testing.T) {
	runner := NewDirect()
	assert.NotNil(t, runner)
	// Mock Action
	mockAction := new(MockAction)
	mockAction.On("Run", nil, "", nil).Return(nil)
	code, data, err := runner.Run(nil, mockAction, "", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, code)
	assert.Equal(t, "mock", data)
}
