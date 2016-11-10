package engine

import (
	"encoding/json"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/engine/runner"
	"github.com/TIBCOSoftware/flogo-lib/util"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("engine")

// Engine creates and executes FlowInstances.
type Engine struct {
	generator      *util.Generator
	runner         action.Runner
	serviceManager *util.ServiceManager
	engineConfig   *Config
	triggersConfig *TriggersConfig
}

// NewEngine create a new Engine
func NewEngine(engineConfig *Config, triggersConfig *TriggersConfig) *Engine {

	var engine Engine
	engine.generator, _ = util.NewGenerator()
	engine.engineConfig = engineConfig
	engine.triggersConfig = triggersConfig
	engine.serviceManager = util.NewServiceManager()

	runnerConfig := engineConfig.RunnerConfig

	if runnerConfig.Type == "direct" {
		engine.runner = runner.NewDirect()
	} else {
		engine.runner = runner.NewPooled(runnerConfig.Pooled)
	}

	if log.IsEnabledFor(logging.DEBUG) {
		cfgJSON, _ := json.MarshalIndent(engineConfig, "", "  ")
		log.Debugf("Engine Configuration:\n%s\n", string(cfgJSON))
	}

	if log.IsEnabledFor(logging.DEBUG) {
		cfgJSON, _ := json.MarshalIndent(triggersConfig, "", "  ")
		log.Debugf("Triggers Configuration:\n%s\n", string(cfgJSON))
	}

	return &engine
}

// RegisterService register a service with the engine
func (e *Engine) RegisterService(service util.Service) {
	e.serviceManager.RegisterService(service)
}

// Start will start the engine, by starting all of its triggers and runner
func (e *Engine) Start() {

	log.Info("Engine: Starting...")

	log.Info("Engine: Starting Services...")

	err := e.serviceManager.Start()

	if err != nil {
		e.serviceManager.Stop()
		panic("Engine: Error Starting Services - " + err.Error())
	}

	log.Info("Engine: Started Services")

	validateTriggers := e.engineConfig.ValidateTriggers

	triggers := trigger.Triggers()

	var triggersToStart []trigger.Trigger

	// initialize triggers
	for _, trigger := range triggers {

		triggerConfig, found := e.triggersConfig.Triggers[trigger.Metadata().ID]

		if !found && validateTriggers {
			panic(fmt.Errorf("Trigger configuration for '%s' not provided", trigger.Metadata().ID))
		}

		if found {
			trigger.Init(triggerConfig, e.runner)
			triggersToStart = append(triggersToStart, trigger)
		}
	}

	runner := e.runner.(interface{})
	managedRunner, ok := runner.(util.Managed)

	if ok {
		util.StartManaged("ActionRunner Service", managedRunner)
	}

	// start triggers
	for _, trigger := range triggersToStart {
		util.StartManaged("Trigger [ "+trigger.Metadata().ID+" ]", trigger)
	}

	log.Info("Engine: Started")
}

// Stop will stop the engine, by stopping all of its triggers and runner
func (e *Engine) Stop() {

	log.Info("Engine: Stopping...")

	triggers := trigger.Triggers()

	// stop triggers
	for _, trigger := range triggers {
		util.StopManaged("Trigger [ "+trigger.Metadata().ID+" ]", trigger)
	}

	runner := e.runner.(interface{})
	managedRunner, ok := runner.(util.Managed)

	if ok {
		util.StopManaged("ActionRunner", managedRunner)
	}

	log.Info("Engine: Stopping Services...")

	err := e.serviceManager.Stop()

	if err != nil {
		log.Error("Engine: Error Stopping Services - " + err.Error())
	} else {
		log.Info("Engine: Stopped Services")
	}

	log.Info("Engine: Stopped")
}
