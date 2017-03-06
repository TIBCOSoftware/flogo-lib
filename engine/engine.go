package engine

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/engine/runner"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/TIBCOSoftware/flogo-lib/util"
)

// Interface for the engine behaviour
// Todo: rename to Engine once the refactoring is completed
type IEngine interface {
	Start()
	Stop()
}

// Engine creates and executes FlowInstances.
type Engine struct {
	generator      *util.Generator
	runner         action.Runner
	serviceManager *util.ServiceManager
	engineConfig   *Config
	triggersConfig *TriggersConfig
}

// EngineConfig is the type for the Engine Configuration
type EngineConfig struct {
	App            *types.AppConfig
	LogLevel       string
	runner         action.Runner
	serviceManager *util.ServiceManager
}

// New creates a new Engine
func New(app *types.AppConfig) (IEngine, error) {
	// App is required
	if app == nil {
		return nil, errors.New("Error: No App configuration provided")
	}
	// Name is required
	if len(app.Name) == 0 {
		return nil, errors.New("Error: No App name provided")
	}
	// Version is required
	if len(app.Version) == 0 {
		return nil, errors.New("Error: No App version provided")
	}

	logLevel := config.GetLogLevel()

	runnerType := config.GetRunnerType()

	var r action.Runner
	// Todo document this values for engine configuration
	if runnerType == "DIRECT" {
		r = runner.NewDirect()
	} else {
		runnerConfig := defaultRunnerConfig()
		r = runner.NewPooled(runnerConfig.Pooled)
	}

	return &EngineConfig{App: app, LogLevel: logLevel, runner: r, serviceManager: util.NewServiceManager()}, nil
}

//Start initializes and starts the Triggers and initializes the Actions
func (e *EngineConfig) Start() {
	logger.Info("Engine: Starting...")

	instanceHelper := app.NewInstanceHelper(e.App, trigger.Factories(), action.Factories())

	// Create the trigger instances
	tInstances, err := instanceHelper.CreateTriggers()
	if err != nil {
		errorMsg := fmt.Sprintf("Engine: Error Creating trigger instances - %s", err.Error())
		logger.Error(errorMsg)
		panic(errorMsg)
	}

	// Initialize and register the triggers
	for key, value := range tInstances {
		triggerConfig := value.Config
		triggerInterface := value.Interf

		//Init
		triggerInterface.Init(*triggerConfig, e.runner)
		//Register
		trigger.RegisterInstance(key, value)
	}

	// Create the action instances
	aInstances, err := instanceHelper.CreateActions()
	if err != nil {
		errorMsg := fmt.Sprintf("Engine: Error Creating action instances - %s", err.Error())
		logger.Error(errorMsg)
		panic(errorMsg)
	}

	// Initialize and register the actions,
	for key, value := range aInstances {
		actionConfig := value.Config
		actionInterface := value.Interf

		//Init
		actionInterface.Init(*actionConfig)
		//Register
		action.RegisterInstance(key, value)

	}

	runner := e.runner.(interface{})
	managedRunner, ok := runner.(util.Managed)

	if ok {
		util.StartManaged("ActionRunner Service", managedRunner)
	}

	// Start the triggers
	for key, value := range tInstances {
		util.StartManaged(fmt.Sprintf("Trigger [ '%s' ]", key), value.Interf)
	}

	logger.Info("Engine: Started")
}

func (e *EngineConfig) Stop() {
	// Todo implement
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

	cfgJSON, _ := json.MarshalIndent(engineConfig, "", "  ")
	logger.Debugf("Engine Configuration:\n%s\n", string(cfgJSON))

	cfgJSON, _ = json.MarshalIndent(triggersConfig, "", "  ")
	logger.Debugf("Triggers Configuration:\n%s\n", string(cfgJSON))

	return &engine
}

// RegisterService register a service with the engine
func (e *Engine) RegisterService(service util.Service) {
	e.serviceManager.RegisterService(service)
}

// Start will start the engine, by starting all of its triggers and runner
func (e *Engine) Start() {

	logger.Info("Engine: Starting...")

	logger.Info("Engine: Starting Services...")

	err := e.serviceManager.Start()

	if err != nil {
		e.serviceManager.Stop()
		panic("Engine: Error Starting Services - " + err.Error())
	}

	logger.Info("Engine: Started Services")

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

	logger.Info("Engine: Started")
}

// Stop will stop the engine, by stopping all of its triggers and runner
func (e *Engine) Stop() {

	logger.Info("Engine: Stopping...")

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

	logger.Info("Engine: Stopping Services...")

	err := e.serviceManager.Stop()

	if err != nil {
		logger.Error("Engine: Error Stopping Services - " + err.Error())
	} else {
		logger.Info("Engine: Stopped Services")
	}

	logger.Info("Engine: Stopped")
}
