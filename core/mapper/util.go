package mapper

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func NewMapperDefFromAnyArray(mappings []interface{}) (*data.MapperDef, error) {

	var mappingDefs []*data.MappingDef

	for _, mapping := range mappings {

		mappingObject := mapping.(map[string]interface{})

		mappingType, err := data.ConvertMappingType(mappingObject["type"])

		if err != nil {
			return nil, err
		}

		value := mappingObject["value"]
		mapTo := mappingObject["mapTo"].(string)

		mappingDef := &data.MappingDef{Type: data.MappingType(mappingType), MapTo: mapTo, Value: value}
		mappingDefs = append(mappingDefs, mappingDef)
	}

	return &data.MapperDef{Mappings: mappingDefs}, nil
}


