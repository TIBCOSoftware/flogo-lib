package json

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/json/field"
	"strconv"
	"strings"

	"sync"

	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("json")

func GetPathValue(value interface{}, refPath string) (interface{}, error) {
	mappingField, err := field.ParseMappingField(refPath)
	if err != nil {
		return nil, fmt.Errorf("parse mapping path [%s] failed, due to %s", refPath, err.Error())
	}

	if mappingField == nil || len(mappingField.Getfields()) <= 0 {
		value, err := makeInterface(value)
		if err != nil {
			value = value
		}
		return value, nil
	}
	return GetFieldValue(value, mappingField)
}

func GetFieldValue(data interface{}, mappingField *field.MappingField) (interface{}, error) {
	var jsonParsed *Container
	var err error
	switch data.(type) {
	case string:
		jsonParsed, err = ParseJSON([]byte(data.(string)))
	default:
		if IsMapperableType(data) {
			jsonParsed, err = Consume(data)
		} else {
			//Take is as string to handle
			b, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			jsonParsed, err = ParseJSON(b)
		}
	}

	if err != nil {
		return nil, err

	}
	return handleGetValue(&JSONData{container: jsonParsed, rw: sync.RWMutex{}}, mappingField.Getfields())
}

func handleGetValue(jsonData *JSONData, fields []string) (interface{}, error) {
	jsonData.rw.Lock()
	defer jsonData.rw.Unlock()

	container := jsonData.container
	if hasArrayFieldInArray(fields) {
		arrayFields, fieldNameindex, arrayIndex := getArrayFieldName(fields)
		//No array field found
		if fieldNameindex == -1 {
			return container.S(arrayFields...).Data(), nil
		}
		restFields := fields[fieldNameindex+1:]
		specialField, err := container.ArrayElement(arrayIndex, arrayFields...)
		if err != nil {
			return nil, err
		}
		log.Debugf("Array element value %s", specialField)
		if hasArrayFieldInArray(restFields) {
			return handleGetValue(&JSONData{container: specialField, rw: sync.RWMutex{}}, restFields)
		}
		return specialField.S(restFields...).Data(), nil
	}
	return container.S(fields...).Data(), nil
}

func getFieldName(fieldName string) string {
	if strings.Index(fieldName, "[") >= 0 {
		return fieldName[0:strings.Index(fieldName, "[")]
	}

	return fieldName
}

func getFieldSliceIndex(fieldName string) (int, error) {
	if strings.Index(fieldName, "[") >= 0 {
		index := fieldName[strings.Index(fieldName, "[")+1 : strings.Index(fieldName, "]")]
		i, err := strconv.Atoi(index)

		if err != nil {
			return -2, nil
		}
		return i, nil
	}

	return -1, nil
}

func getNameInsideBrancket(fieldName string) string {
	if strings.Index(fieldName, "[") >= 0 {
		index := fieldName[strings.Index(fieldName, "[")+1 : strings.Index(fieldName, "]")]
		return index
	}

	return ""
}

func makeInterface(value interface{}) (interface{}, error) {

	var paramMap interface{}

	if value == nil {
		return paramMap, nil
	}

	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(t), &paramMap)
		if err != nil {
			return nil, err
		}
		return paramMap, nil
	default:
		return value, nil
	}
	return paramMap, nil
}

func getArrayFieldName(fields []string) ([]string, int, int) {
	var tmpFields []string
	index := -1
	var arrayIndex int
	for i, field := range fields {
		if strings.Index(field, "[") >= 0 && strings.Index(field, "]") >= 0 {
			arrayIndex, _ = getFieldSliceIndex(field)
			fieldName := getFieldName(field)
			index = i
			if fieldName != "" {
				tmpFields = append(tmpFields, getFieldName(field))
			}
			break
		} else {
			tmpFields = append(tmpFields, field)
		}
	}
	return tmpFields, index, arrayIndex
}

func hasArrayFieldInArray(fields []string) bool {
	for _, field := range fields {
		if strings.Index(field, "[") >= 0 && strings.HasSuffix(field, "]") {
			//Make sure the index are integer
			_, err := strconv.Atoi(getNameInsideBrancket(field))
			if err == nil {
				return true
			}
		}
	}
	return false
}
