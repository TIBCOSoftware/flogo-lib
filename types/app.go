package types

import (
	"encoding/json"
)

// App is the configuration for the App
type AppConfig struct {
	Name        string           `json:"name"`
	Version     string           `json:"version"`
	Description string           `json:"description"`
	Triggers    []*TriggerConfig `json:"triggers"`
	Actions     []*ActionConfig  `json:"actions"`
}

// Trigger is the configuration for the Trigger
type TriggerConfig struct {
	Id       string                 `json:"id"`
	Ref      string                 `json:"ref"`
	Settings map[string]interface{} `json:"settings"`
	Handlers []*TriggerHandler      `json:"handlers"`
}

// TriggerHandler is the configuration for the Trigger Handler
type TriggerHandler struct {
	ActionId string                 `json:"actionId"`
	Settings map[string]interface{} `json:"settings"`
}

// Action is the configuration for the Action
type ActionConfig struct {
	Id   string          `json:"id"`
	Ref  string          `json:"ref"`
	Data json.RawMessage `json:"data"`
}
