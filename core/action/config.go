package action

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

// Config is the configuration for the Action
type Config struct {
	//inline action
	Ref      string           `json:"ref"`
	Data     json.RawMessage  `json:"data"`

	//referenced action
	Id       string           `json:"id"`

	// Deprecated: No longer used
	Metadata *data.IOMetadata `json:"metadata"`
}
