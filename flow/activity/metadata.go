package activity

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

// Metadata is the metadata for the Activity
type Metadata struct {
	ID       string
	Settings map[string]*data.Attribute
	Inputs   map[string]*data.Attribute
	Outputs  map[string]*data.Attribute
}

// NewMetadata creates the metadata object from its json representation
func NewMetadata(jsonMetadata string) *Metadata {
	md := &Metadata{}
	err := json.Unmarshal([]byte(jsonMetadata), md)
	if err != nil {
		panic("Unable to parse activity metadata: " + err.Error())
	}

	return md
}

// MarshalJSON overrides the default MarshalJSON for TaskEnv
func (md *Metadata) MarshalJSON() ([]byte, error) {

	inputs := make([]*data.Attribute, 0, len(md.Inputs))

	for _, value := range md.Inputs {
		inputs = append(inputs, value)
	}

	outputs := make([]*data.Attribute, 0, len(md.Outputs))

	for _, value := range md.Outputs {
		outputs = append(outputs, value)
	}

	settings := make([]*data.Attribute, 0, len(md.Settings))
	if md.Settings != nil {
		for _, value := range md.Settings {
			settings = append(settings, value)
		}
	}

	return json.Marshal(&struct {
		Name     string            `json:"name"`
		Settings []*data.Attribute `json:"settings"`
		Inputs   []*data.Attribute `json:"inputs"`
		Outputs  []*data.Attribute `json:"outputs"`
	}{
		Name:     md.ID,
		Settings: settings,
		Inputs:   inputs,
		Outputs:  outputs,
	})
}

// UnmarshalJSON overrides the default UnmarshalJSON for TaskEnv
func (md *Metadata) UnmarshalJSON(b []byte) error {

	ser := &struct {
		Name    string            `json:"name"`
		Settings []*data.Attribute `json:"settings"`
		Inputs  []*data.Attribute `json:"inputs"`
		Outputs []*data.Attribute `json:"outputs"`
	}{}

	if err := json.Unmarshal(b, ser); err != nil {
		return err
	}

	md.ID = ser.Name
	md.Inputs = make(map[string]*data.Attribute, len(ser.Inputs))
	md.Outputs = make(map[string]*data.Attribute, len(ser.Outputs))
	md.Settings = make(map[string]*data.Attribute, len(ser.Settings))

	for _, attr := range ser.Inputs {
		md.Inputs[attr.Name] = attr
	}

	for _, attr := range ser.Outputs {
		md.Outputs[attr.Name] = attr
	}
	
	for _, attr := range ser.Settings {
		md.Settings[attr.Name] = attr
	}

	return nil
}
