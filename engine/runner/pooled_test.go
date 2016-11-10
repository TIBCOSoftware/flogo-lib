package runner

import (
	"context"
	"errors"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFullAction struct {
	mock.Mock
}

func (m *MockFullAction) Run(context context.Context, uri string, options interface{}, handler action.ResultHandler) error {
	args := m.Called(context, uri, options, handler)
	return args.Error(0)
}

// This mock action will handle the result and mark it done
type MockResultAction struct {
	mock.Mock
}

func (m *MockResultAction) Run(context context.Context, uri string, options interface{}, handler action.ResultHandler) error {
	args := m.Called(context, uri, options, handler)
	go func() {
		handler.HandleResult(0, "mock", nil)
		handler.Done()
	}()
	return args.Error(0)
}

// TestNewPooledOk test creation of new Pooled runner
func TestNewPooledOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 1, WorkQueueSize: 1}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
}

// TestStartOk test that Start method is fine
func TestStartOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 3, WorkQueueSize: 3}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	// It should have a worker queue of the size expected
	assert.Equal(t, 3, cap(runner.workerQueue))
	// It should have a workers of the expected size
	assert.Equal(t, 3, len(runner.workers))
	// Runner should be active
	assert.True(t, runner.active)
}

// TestRunNilError test that running a nil action trows and error
func TestRunNilError(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	_, _, err = runner.Run(nil, nil, "", nil)
	assert.NotNil(t, err)
}

// TestRunInnactiveError test that running an innactive runner trows and error
func TestRunInnactiveError(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	a := new(MockFullAction)
	_, _, err := runner.Run(nil, a, "", nil)
	assert.NotNil(t, err)
}

// TestRunErrorInAction test that running an action returns an error
func TestRunErrorInAction(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	a := new(MockFullAction)
	a.On("Run", nil, mock.AnythingOfType("string"), nil, mock.AnythingOfType("*runner.AsyncResultHandler")).Return(errors.New("Error in action"))
	_, _, err = runner.Run(nil, a, "mockAction", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Error in action", err.Error())
}

// TestRunOk test that running an action is ok
func TestRunOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	a := new(MockResultAction)
	a.On("Run", nil, mock.AnythingOfType("string"), nil, mock.AnythingOfType("*runner.AsyncResultHandler")).Return(nil)
	code, data, err := runner.Run(nil, a, "mockAction", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, code)
	assert.Equal(t, "mock", data)
}

// TestStopOk test that Stop method is fine
func TestStopOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 3, WorkQueueSize: 3}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	err = runner.Stop()
	assert.Nil(t, err)
	assert.False(t, runner.active)

}
