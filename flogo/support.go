package flogo

import (
	"context"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

func toAppConfig(a *App) *app.Config {

	appCfg := &app.Config{}
	appCfg.Name = "app"
	appCfg.Version = "1.0.0"
	appCfg.Properties = a.Properties()
	appCfg.Resources = a.resources

	var triggerConfigs []*trigger.Config
	for _, trg := range a.Triggers() {

		triggerConfigs = append(triggerConfigs, toTriggerConfig(trg))
	}

	appCfg.Triggers = triggerConfigs

	return appCfg
}

func toTriggerConfig(trg *Trigger) *trigger.Config {

	triggerConfig := &trigger.Config{Ref: trg.ref, Settings: trg.Settings()}

	//todo add output
	//trigger.Config struct { Output   map[string]interface{} `json:"output"` }

	var handlerConfigs []*trigger.HandlerConfig
	for _, handler := range trg.Handlers() {
		h := &trigger.HandlerConfig{Settings: handler.Settings()}
		//todo add output
		//trigger.HandlerConfig struct { Output   map[string]interface{} `json:"output"` }

		//todo only handles one action for now
		for _, act := range handler.Actions() {
			h.Action = toActionConfig(act)
			break
		}

		handlerConfigs = append(handlerConfigs, h)
	}

	triggerConfig.Handlers = handlerConfigs
	return triggerConfig
}

func toActionConfig(act *Action) *action.Config {
	actionCfg := &action.Config{}

	if act.act != nil {
		actionCfg.Act = act.act
		return actionCfg
	}

	actionCfg.Ref = act.ref
	// convert settings to "data"
	// action.Config struct { Data json.RawMessage  `json:"data"` }
	mappings := &data.IOMappings{}

	if len(act.inputMappings) > 0 {
		mappings.Input, _ = toMappingDefs(act.inputMappings)
	}
	if len(act.outputMappings) > 0 {
		mappings.Output, _ = toMappingDefs(act.outputMappings)
	}
	actionCfg.Mappings = mappings

	return actionCfg
}

func toMappingDefs(mappings []string) ([]*data.MappingDef, error) {

	var mappingDefs []*data.MappingDef
	for _, strMapping := range mappings {

		idx := strings.Index(strMapping, "=")
		lhs := strings.TrimSpace(strMapping[:idx])
		rhs := strings.TrimSpace(strMapping[idx+1:])

		mType, mValue := getMappingValue(rhs)
		mappingDef := &data.MappingDef{Type: mType, MapTo: lhs, Value: mValue}
		mappingDefs = append(mappingDefs, mappingDef)
	}
	return mappingDefs, nil
}

func getMappingValue(strValue string) (data.MappingType, interface{}) {

	//todo add support for other mapping types
	return data.MtExpression, strValue
}

type proxyAction struct {
	handlerFunc HandlerFunc
	md *action.Metadata
}

func NewProxyAction(f HandlerFunc) action.Action {
	return &proxyAction{
		handlerFunc: f,
		md: &action.Metadata{Async:false},
	}
}

// Metadata get the Action's metadata
func (a *proxyAction)Metadata() *action.Metadata {
	return a.md
}

// IOMetadata get the Action's IO metadata
func (a *proxyAction) IOMetadata() *data.IOMetadata {
	return nil
}

func (a *proxyAction) Run(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {
	return a.handlerFunc(ctx, inputs)
}
