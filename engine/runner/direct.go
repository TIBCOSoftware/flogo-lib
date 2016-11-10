package runner

import (
	"context"
	"errors"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("runner")

// DirectRunner runs an action synchronously
type DirectRunner struct {
}

// NewDirectRunner create a new DirectRunner
func NewDirect() *DirectRunner {
	return &DirectRunner{}
}

// Start will start the engine, by starting all of its workers
func (runner *DirectRunner) Start() error {
	//op-op
	return nil
}

// Stop will stop the engine, by stopping all of its workers
func (runner *DirectRunner) Stop() error {
	//no-op
	return nil
}

// Run the specified action
func (runner *DirectRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {

	if action == nil {
		return 0, nil, errors.New("Action not found")
	}

	handler := &SyncResultHandler{done: make(chan bool, 1)}

	err = action.Run(context, uri, options, handler)

	if err != nil {
		return 0, nil, err
	}

	<-handler.done

	return handler.Result()
}

// SyncResultHandler simple result handler to use in synchronous case
type SyncResultHandler struct {
	done chan (bool)
	code int
	data interface{}
	err  error
}

// HandleResult implements action.ResultHandler.HandleResult
func (rh *SyncResultHandler) HandleResult(code int, data interface{}, err error) {
	rh.code = code
	rh.data = data
	rh.err = err
}

// Done implements action.ResultHandler.Done
func (rh *SyncResultHandler) Done() {
	rh.done <- true
}

// Result returns the latest Result set on the handler
func (rh *SyncResultHandler) Result() (code int, data interface{}, err error) {
	return rh.code, rh.data, rh.err
}
