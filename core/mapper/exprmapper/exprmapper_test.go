package exprmapper

import (
	"testing"

	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/stretchr/testify/assert"
)

//1. activity mapping
func TestActivityMapping(t *testing.T) {
	v, err := expressMap("$activity[a1].field.id", "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "d", v)
}

//Literal string
func TestMappingRef(t *testing.T) {
	v, err := expressMap("dddd", "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "dddd", v)
}

//2. flow mapping
func TestFlowMapping(t *testing.T) {
	v, err := expressMapWithFlow("$.field.id", "field", data.GetBasicResolver())
	assert.Nil(t, err)
	assert.Equal(t, "d", v)
}

//3. function
func TestGetMapValueFunction(t *testing.T) {
	v, err := expressMap(`string.concat("ddddd",$activity[a1].field.id)`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "dddddd", v)
	v, err = expressMap(`string.concat("ddddd",$activity[a1].field.id, string.concat($activity[a1].field.id,$activity[a1].field.id))`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "dddddddd", v)
}

//4. function with different type
func TestGetMapValueFunctionTypes(t *testing.T) {

	v, err := expressMap(`string.concat(123,$activity[a1].field.id)`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "123d", v)

	v, err = expressMap(`string.concat("s-",$activity[a1].field.id, string.concat($activity[a1].field.id,$activity[a1].field.id),true)`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "s-dddtrue", v)

	v, err = expressMap(`string.concat(123,$activity[a1].field.id, "dddd", 450)`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "123ddddd450", v)
}

//5. expression
//6. ternary expression
func TestGetMapValueExpression(t *testing.T) {

	v, err := expressMapBoolMapToField(`string.length(string.concat("ddddd",$activity[a1].field.id)) == 6`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	v, err = expressMapBoolMapToField(`$activity[a1].field.id == "d"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	v, err = expressMapBoolMapToField(`($activity[a1].field.id == "d") == true`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	v, err = expressMap(`$activity[a1].field.id == "d" ? $activity[a1].field.id : string.concat("ssss",$activity[a1].field.id)`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "d", v)

	v, err = expressMap(`$activity[a1].field.id == "d" ? string.concat("ssss",$activity[a1].field.id) : $activity[a1].field.id`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "ssssd", v)

	v, err = expressMap(`$activity[a1].field.id != "d" ? string.concat("ssss",$activity[a1].field.id) : $activity[a1].field.id`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "d", v)

	v, err = expressMap(`$activity[a1].field.id != "d" ? "dddd":"ssss"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "ssss", v)

	v, err = expressMap(`$activity[a1].field.id == "d" ? "dddd":"ssss"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "dddd", v)

	v, err = expressMap(` 2>1 ? "dddd":"ssss"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "dddd", v)

	v, err = expressMap(` 2<1 ? "dddd":"ssss"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "ssss", v)

}

func TestOperatorPriority(t *testing.T) {

	v, err := expressMap(`string.concat(123,$activity[a1].field.id) == "123d" && $activity[a1].field.id =="d"`, "field", GetTestResolver())
	assert.Nil(t, err)
	assert.Equal(t, "true", v)

	v, err = expressMap(`string.concat(123,$activity[a1].field.id) == "123d" && $activity[a1].field.id !="d"`, "field", GetTestResolver())
	assert.Nil(t, err)

	assert.Equal(t, "false", v)

}

func getSimpleScope(name, value string, fieldType data.Type) data.Scope {
	a, _ := data.NewAttribute(name, fieldType, value)
	maps := make(map[string]*data.Attribute)
	maps[name] = a
	scope := data.NewFixedScope(maps)
	scope.SetAttrValue(name, value)
	return scope
}

func GetObjectFieldScope(name, value string) data.Scope {
	return getSimpleScope(name, value, data.TypeObject)
}

func GetStringFieldScope(name, value string) data.Scope {
	return getSimpleScope(name, value, data.TypeString)
}

func GetBoolFieldScope(name, value string) data.Scope {
	return getSimpleScope(name, value, data.TypeBoolean)
}

func GetComplexxFieldScope(name, value string) data.Scope {
	return getSimpleScope(name, value, data.TypeComplexObject)
}

//For test  purpuse copy it from flogo-contrib
var resolver = &TestResolver{}

func GetTestResolver() data.Resolver {
	return resolver
}

type TestResolver struct {
}

func (r *TestResolver) Resolve(toResolve string, scope data.Scope) (value interface{}, err error) {

	var details *data.ResolutionDetails

	if strings.HasPrefix(toResolve, "${") {
		details, err = data.GetResolutionDetailsOld(toResolve)
	} else if strings.HasPrefix(toResolve, "$") {
		details, err = data.GetResolutionDetails(toResolve[1:])
	} else {
		return data.SimpleScopeResolve(toResolve, scope)
	}

	if err != nil {
		return nil, err
	}

	if details == nil {
		return nil, fmt.Errorf("unable to determine resolver for %s", toResolve)
	}

	switch details.ResolverName {
	case "activity":
		attr, exists := scope.GetAttr("_A." + details.Item + "." + details.Property)
		if !exists {
			return nil, fmt.Errorf("failed to resolve activity attr: '%s', not found in flow", details.Property)
		}
		value = attr.Value()
	default:
		return nil, fmt.Errorf("unsupported resolver: %s", details.ResolverName)
	}

	if details.Path != "" {
		value, err = data.PathGetValue(value, details.Path)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return value, nil
}

func expressMap(from, to string, resolver data.Resolver) (interface{}, error) {
	mapDef := &data.MappingDef{Type: data.MtExpression, Value: from, MapTo: to}
	inputScope := GetObjectFieldScope("_A.a1.field", `{"id":"d"}`)
	outputScope := GetStringFieldScope("field", "")
	err := MapExpreesion(mapDef, inputScope, outputScope, resolver)
	if err != nil {
		return nil, err
	}
	arr, ok := outputScope.GetAttr("field")
	if !ok {
		return nil, fmt.Errorf("Cannot find attribute [%s] in output scope", "field")
	}
	return arr.Value(), nil
}

func expressMapBoolMapToField(from, to string, resolver data.Resolver) (interface{}, error) {
	mapDef := &data.MappingDef{Type: data.MtExpression, Value: from, MapTo: to}
	inputScope := GetObjectFieldScope("_A.a1.field", `{"id":"d"}`)
	outputScope := GetBoolFieldScope("field", "")
	err := MapExpreesion(mapDef, inputScope, outputScope, resolver)
	if err != nil {
		return nil, err
	}
	arr, ok := outputScope.GetAttr("field")
	if !ok {
		return nil, fmt.Errorf("Cannot find attribute [%s] in output scope", "field")
	}
	return arr.Value(), nil
}

func expressMapWithFlow(from, to string, resolver data.Resolver) (interface{}, error) {
	mapDef := &data.MappingDef{Type: data.MtExpression, Value: from, MapTo: to}
	inputScope := GetObjectFieldScope("field", `{"id":"d"}`)
	outputScope := GetStringFieldScope("field", "")
	err := MapExpreesion(mapDef, inputScope, outputScope, resolver)
	if err != nil {
		return nil, err
	}
	arr, ok := outputScope.GetAttr("field")
	if !ok {
		return nil, fmt.Errorf("Cannot find attribute [%s] in output scope", "field")
	}
	return arr.Value(), nil
}
