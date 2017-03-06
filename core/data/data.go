package data

import (
	"encoding/json"
)

// Attribute is a simple structure used to define a data Attribute/property
type Attribute struct {
	Name  string
	Type  Type
	Value interface{}
}

// NewAttribute constructs a new attribute
func NewAttribute(name string, attrType Type, value interface{}) *Attribute {
	var attr Attribute
	attr.Name = name
	attr.Type = attrType
	attr.Value = value

	return &attr
}

// MarshalJSON implements json.Marshaler.MarshalJSON
func (tv *Attribute) MarshalJSON() ([]byte, error) {

	return json.Marshal(&struct {
		Name  string      `json:"name"`
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}{
		Name:  tv.Name,
		Type:  tv.Type.String(),
		Value: tv.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler.UnmarshalJSON
func (tv *Attribute) UnmarshalJSON(data []byte) error {

	ser := &struct {
		Name  string      `json:"name"`
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}{}

	if err := json.Unmarshal(data, ser); err != nil {
		return err
	}

	tv.Name = ser.Name
	tv.Type, _ = ToTypeEnum(ser.Type)
	val, err := CoerceToValue(ser.Value, tv.Type)
	if err != nil {
		//todo what should we do if there is an error coercing the value?
		tv.Value = ser.Value
	} else {
		tv.Value = val
	}

	return nil
}

// ComplexObject is the value that is used when using a "COMPLEX_OBJECT" type
type ComplexObject struct {
	Metadata string `json:"metadata"`
	Value    interface{} `json:"value"`
}