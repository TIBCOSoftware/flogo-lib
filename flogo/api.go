package flogo

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/TIBCOSoftware/flogo-lib/app/resource"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/engine"
)

type App struct {
	properties []*data.Attribute
	triggers   []*Trigger
	resources  []*resource.Config
}

type Trigger struct {
	ref      string
	settings map[string]interface{}
	handlers []*Handler
}

type Handler struct {
	settings map[string]interface{}
	actions  []*Action
}

type HandlerFunc func(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error)

type Action struct {
	ref            string
	act            action.Action
	settings       map[string]interface{}
	inputMappings  []string
	outputMappings []string
}

func NewApp() *App {
	return &App{}
}

func (a *App) NewTrigger(trg trigger.Trigger, settings map[string]interface{}) *Trigger {

	value := reflect.ValueOf(trg)
	value = value.Elem()
	ref := value.Type().PkgPath()

	newTrg := &Trigger{ref: ref, settings: settings}
	a.triggers = append(a.triggers, newTrg)

	return newTrg
}

func (a *App) AddProperty(name string, dataType data.Type, value interface{}) error {
	property, err := data.NewAttribute(name, dataType, value)
	if err != nil {
		return err
	}
	a.properties = append(a.properties, property)
	return nil
}

func (a *App) AddResource(id string, data json.RawMessage) {

	res := &resource.Config{ID: id, Data: data}
	a.resources = append(a.resources, res)
}

func (a *App) Properties() []*data.Attribute {

	return a.properties
}

func (a *App) Triggers() []*Trigger {

	return a.triggers
}

func (t *Trigger) Settings() map[string]interface{} {

	return t.settings
}

func (t *Trigger) NewHandler(settings map[string]interface{}) *Handler {

	newHandler := &Handler{settings: settings}
	t.handlers = append(t.handlers, newHandler)

	return newHandler
}

func (t *Trigger) NewFuncHandler(settings map[string]interface{}, handlerFunc HandlerFunc) *Handler {

	newHandler := &Handler{settings: settings}
	newAct := &Action{act: NewProxyAction(handlerFunc)}
	newHandler.actions = append(newHandler.actions, newAct)

	t.handlers = append(t.handlers, newHandler)

	return newHandler
}

func (t *Trigger) Handlers() []*Handler {
	return t.handlers
}

func (h *Handler) Settings() map[string]interface{} {
	return h.settings
}

func (h *Handler) NewAction(act action.Action, settings map[string]interface{}) *Action {

	value := reflect.ValueOf(act)
	value = value.Elem()
	ref := value.Type().PkgPath()

	newAct := &Action{ref: ref, settings: settings}
	h.actions = append(h.actions, newAct)

	return newAct
}

func (h *Handler) Actions() []*Action {
	return h.actions
}

func (a *Action) Settings() map[string]interface{} {
	return a.settings
}

func (a *Action) SetInputMappings(mappings ...string) {
	a.inputMappings = mappings
}

func (a *Action) SetOutputMappings(mappings ...string) {
	a.outputMappings = mappings
}

func (a *Action) InputMappings() []string {
	return a.inputMappings
}

func (a *Action) OutputMappings() []string {
	return a.outputMappings
}

func NewEngine(a *App) (engine.Engine, error) {
	appConfig := toAppConfig(a)
	return engine.New(appConfig)
}
