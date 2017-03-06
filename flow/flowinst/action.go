package flowinst

import (
	"context"
	"errors"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/flow/flowdef"
	"github.com/TIBCOSoftware/flogo-lib/util"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	AoStart   = iota // 0
	AoResume         // 1
	AoRestart        // 2
)

// ActionType the name/type of the FlowAction
const ActionType = "flow"

// ActionOptions are the options for the FlowAction
type ActionOptions struct {
	MaxStepCount int
	Record       bool
}

// FlowAction is a Action that executes a flow
type FlowAction struct {
	stateRecorder StateRecorder
	flowProvider  flowdef.Provider
	idGenerator   *util.Generator
	actionOptions *ActionOptions
}

// NewFlowAction creates a new FlowAction
func NewFlowAction(flowProvider flowdef.Provider, stateRecorder StateRecorder, options *ActionOptions) *FlowAction {
	var action FlowAction
	action.flowProvider = flowProvider
	action.stateRecorder = stateRecorder
	action.idGenerator, _ = util.NewGenerator()
	// fix up run options

	if options == nil {
		options = &ActionOptions{Record: true}
	}

	if options.MaxStepCount < 1 {
		options.MaxStepCount = int(^uint16(0))
	}

	options.Record = (stateRecorder != nil) && options.Record

	action.actionOptions = options

	return &action
}

// RunOptions the options when running a FlowAction
type RunOptions struct {
	Op           int
	ReturnID     bool
	InitialState *Instance
	ExecOptions  *ExecOptions
}

// Run implements action.Action.Run
func (fa *FlowAction) Run(context context.Context, uri string, options interface{}, handler action.ResultHandler) error {

	//todo: catch panic
	//todo: consider switch to URI to dictate flow operation (ex. flow://blah/resume)

	op := AoStart
	retID := false

	ro, ok := options.(*RunOptions)

	if ok {
		op = ro.Op
		retID = ro.ReturnID
	}

	var instance *Instance

	switch op {
	case AoStart:
		flow, _ := fa.flowProvider.GetFlow(uri)

		if flow == nil {
			err := fmt.Errorf("Flow [%s] not found", uri)
			return err
		}

		instanceID := fa.idGenerator.NextAsString()
		logger.Debug("Creating Instance: ", instanceID)

		instance = NewFlowInstance(instanceID, uri, flow)
	case AoResume:
		if ok {
			instance = ro.InitialState
			logger.Debug("Resuming Instance: ", instance.ID())
		} else {
			return errors.New("Unable to resume instance, resume options not provided")
		}
	case AoRestart:
		if ok {
			instance = ro.InitialState
			instanceID := fa.idGenerator.NextAsString()
			instance.Restart(instanceID, fa.flowProvider)

			logger.Debug("Restarting Instance: ", instanceID)
		} else {
			return errors.New("Unable to restart instance, restart options not provided")
		}
	}

	if ok && ro.ExecOptions != nil {
		logger.Debugf("Applying Exec Options to instance: %s\n", instance.ID())
		ApplyExecOptions(instance, ro.ExecOptions)
	}

	triggerAttrs, ok := trigger.FromContext(context)

	if ok {
		if len(triggerAttrs) > 0 {
			logger.Debug("Run Attributes:")
			for _, attr := range triggerAttrs {
				logger.Debugf(" Attr:%s, Type:%s, Value:%v", attr.Name, attr.Type.String(), attr.Value)
			}
		}
	}

	if op == AoStart {
		instance.Start(triggerAttrs)
	} else {
		instance.UpdateAttrs(triggerAttrs)
	}

	logger.Debugf("Executing instance: %s\n", instance.ID())

	stepCount := 0
	hasWork := true

	instance.SetReplyHandler(&SimpleReplyHandler{resultHandler: handler})

	go func() {

		defer handler.Done()

		if !instance.Flow.ExplicitReply() {
			handler.HandleResult(200, &IDResponse{ID: instance.ID()}, nil)
		}

		for hasWork && instance.Status() < StatusCompleted && stepCount < fa.actionOptions.MaxStepCount {
			stepCount++
			logger.Debugf("Step: %d\n", stepCount)
			hasWork = instance.DoStep()

			if fa.actionOptions.Record {
				fa.stateRecorder.RecordSnapshot(instance)
				fa.stateRecorder.RecordStep(instance)
			}
		}

		if retID {
			handler.HandleResult(200, &IDResponse{ID: instance.ID()}, nil)
		}

		logger.Debugf("Done Executing A.instance [%s] - Status: %d\n", instance.ID(), instance.Status())

		if instance.Status() == StatusCompleted {
			logger.Infof("Flow [%s] Completed", instance.ID())
		}
	}()

	return nil
}

// SimpleReplyHandler is a simple ReplyHandler that is pass-thru to the action ResultHandler
type SimpleReplyHandler struct {
	resultHandler action.ResultHandler
}

// Reply implements ReplyHandler.Reply
func (rh *SimpleReplyHandler) Reply(replyCode int, replyData interface{}, err error) {

	rh.resultHandler.HandleResult(replyCode, replyData, err)
}

// IDResponse is a response object consists of an ID
type IDResponse struct {
	ID string `json:"id"`
}
