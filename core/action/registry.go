package action

import (
	"fmt"
	"sync"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var (
	actionsMu sync.Mutex
	actions   = make(map[string]Action)
	reg       = &registry{}
)

type Registry interface {
	AddFactory(ref string, f Factory) error
	GetFactories() map[string]Factory
	AddInstance(id string, instance *ActionInstance) error
	GetAction(id string) Action2
}

type registry struct {
	factories map[string]Factory
	instances map[string]*ActionInstance
}

func RegisterFactory(ref string, f Factory) error {
	return reg.AddFactory(ref, f)
}

func (r *registry) AddFactory(ref string, f Factory) error {
	actionsMu.Lock()
	defer actionsMu.Unlock()

	logger.Debugf("Registering action factory ref: '%s'", ref)

	if len(ref) == 0 {
		return fmt.Errorf("registry.RegisterFactory: ref is empty")
	}

	if f == nil {
		return fmt.Errorf("registry.RegisterFactory: factory is nil")
	}

	// copy on write to avoid synchronization on access
	newFs := make(map[string]Factory, len(r.factories))

	for k, v := range r.factories {
		newFs[k] = v
	}

	if newFs[ref] != nil {
		return fmt.Errorf("registry.RegisterFactory: already registered factory for ref '%s'", ref)
	}

	newFs[ref] = f

	r.factories = newFs

	return nil
}

func Factories() map[string]Factory {
	return reg.GetFactories()
}

// GetFactories returns a copy of the factories map
func (r *registry) GetFactories() map[string]Factory {

	newFs := make(map[string]Factory, len(r.factories))

	for k, v := range r.factories {
		newFs[k] = v
	}

	return newFs
}

func RegisterInstance(id string, inst *ActionInstance) error {
	return reg.AddInstance(id, inst)
}

func (r *registry) AddInstance(id string, inst *ActionInstance) error {
	actionsMu.Lock()
	defer actionsMu.Unlock()

	if len(id) == 0 {
		return fmt.Errorf("registry.RegisterInstance: id is empty")
	}

	if inst == nil {
		return fmt.Errorf("registry.RegisterInstance: instance is nil")
	}

	// copy on write to avoid synchronization on access
	newInst := make(map[string]*ActionInstance, len(r.instances))

	for k, v := range r.instances {
		newInst[k] = v
	}

	if newInst[id] != nil {
		return fmt.Errorf("registry.RegisterInstance: already registered instance for id '%s'", id)
	}

	newInst[id] = inst

	r.instances = newInst

	return nil
}

// Register registers the specified action
func Register(actionType string, action Action) {
	actionsMu.Lock()
	defer actionsMu.Unlock()

	if actionType == "" {
		panic("action.Register: actionType is empty")
	}

	if action == nil {
		panic("action.Register: action is nil")
	}

	if _, dup := actions[actionType]; dup {
		panic("action.Register: action already registered for action type: " + actionType)
	}

	// copy on write to avoid synchronization on access
	newActions := make(map[string]Action, len(actions))

	for k, v := range actions {
		newActions[k] = v
	}

	newActions[actionType] = action
	actions = newActions

	logger.Debugf("Registerd Action: %s", actionType)
}

// Actions gets all the registered Action Actions
func Actions() []Action {

	var curActions = actions

	list := make([]Action, 0, len(curActions))

	for _, value := range curActions {
		list = append(list, value)
	}

	return list
}

// Get gets specified Action
func Get(actionType string) Action {
	return actions[actionType]
}

// Get gets specified Action
func Get2(id string) Action2 {
	return reg.GetAction(id)
}

// Get gets specified Action
func (r *registry) GetAction(id string) Action2 {
	instance := r.instances[id]
	if instance != nil {
		return instance.Interf
	}
	return nil
}
