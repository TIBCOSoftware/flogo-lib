package expression

import (
	"encoding/json"
	"fmt"
	"testing"

	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
)

func TestExpressionTernary(t *testing.T) {
	v, err := ParseExpression(`1>2?string.concat("sss","ddddd"):"fff"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	s, _ := json.Marshal(v)
	fmt.Println("-------------------", string(s))
	result, err := v.EvalWithScope(nil, nil)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "fff", result)

	fmt.Println("Result:", result)
}

func TestExpressionTernaryString(t *testing.T) {
	v, err := ParseExpression(`1<2?"lixingwang":"fff"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	result, err := v.EvalWithScope(nil, nil)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "lixingwang", result)
	fmt.Println("Result:", result)
}

func TestExpressionTernaryString3(t *testing.T) {
	v, err := ParseExpression(`200>100?true:false`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	result, err := v.EvalWithScope(nil, nil)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, result)
	fmt.Println("Result:", result)
}

func TestExpressionString(t *testing.T) {
	v, err := ParseExpression(`$activity[C].result==3`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	_, err = v.EvalWithScope(nil, nil)
	assert.NotNil(t, err)
}

func TestExpressionWithOldWay(t *testing.T) {
	v, err := ParseExpression(`${flow.petMax} > 2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	_, err = v.EvalWithScope(nil, nil)
	assert.NotNil(t, err)

}

func TestExpressionTernaryFunction(t *testing.T) {
	v, err := ParseExpression(`string.length($TriggerData.queryParams.id) == 0 ? "Query Id cannot be null" : string.length($TriggerData.queryParams.id)`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	//a, _ := data.NewAttribute("queryParams", data.TypeComplexObject, &data.ComplexObject{Metadata: "", Value: `{"id":"lixingwang"}`})
	//metadata := make(map[string]*data.Attribute)
	//metadata["queryParams"] = a

	//scope.SetAttrValue("queryParams", &data.ComplexObject{Metadata: "", Value: `{"id":"lixingwang"}`})
	result, err := v.EvalWithScope(nil, nil)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, _ := json.Marshal(result)
	//assert.Equal(t, "lixingwang", result)
	fmt.Println("Result:", string(vv))
}

func TestExpressionTernaryRef(t *testing.T) {
	os.Setenv("name", "flogo")
	os.Setenv("address", "tibco")

	v, err := ParseExpression(`string.length("lixingwang")>11?$env.name:$env.address`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	result, err := v.EvalWithScope(data.NewFixedScope(nil), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	assert.Equal(t, "tibco", result)

	fmt.Println("Result:", result)
}

func TestExpressionTernaryRef2(t *testing.T) {
	v, err := ParseExpression(`string.length("lixingwang")>11?"lixingwang":"fff"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	s, _ := json.Marshal(v)
	fmt.Println("-------------------", string(s))
	result, err := v.EvalWithScope(nil, nil)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "fff", result)

	fmt.Println("Result:", result)
}

func TestWeExpr_LinkMapping(t *testing.T) {
	expr, err := ParseExpression(`$T.parameters.path_params[0].value==2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", expr)
}

func TestWeExpr_LinkMapping2(t *testing.T) {
	v, err := ParseExpression(`$T.parameters==2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
}

func TestExpressionInt(t *testing.T) {
	expr, err := ParseExpression(`123==456`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := expr.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, false, v)

	fmt.Println("Result:", v)
}

func TestExpressionEQ(t *testing.T) {
	expr, err := ParseExpression(`123==123`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := expr.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	fmt.Println("Result:", v)
}

func TestExpressionEQFunction(t *testing.T) {
	expr, err := ParseExpression(`string.concat("123","456")=="123456"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := expr.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
	fmt.Println("Result:", v)
}

func TestExpressionEQFunction2Side(t *testing.T) {
	e, err := ParseExpression(`string.concat("123","456") == string.concat("12","3456")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
	fmt.Println("Result:", v)
}

func TestExpressionRef(t *testing.T) {
	_, err := ParseExpression(`$A4.query.name=="name"`)
	assert.Nil(t, err)
}

func TestExpressionFunction(t *testing.T) {
	e, err := ParseExpression(`string.concat("tibco","software")=="tibcosoftware"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	fmt.Println("Result:", v)
}

func TestExpressionAnd(t *testing.T) {
	e, err := ParseExpression(`("dddddd" == "dddd3dd") && ("133" == "123")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, false, v)
	fmt.Println("Result:", v)
}

func TestExpressionOr(t *testing.T) {
	e, err := ParseExpression(`("dddddd" == "dddddd") && ("123" == "123")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
	fmt.Println("Result:", v)
}

func TestFunc(t *testing.T) {
	e, err := ParseExpression(`string.length("lixingwang") == 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, true)
	e, err = ParseExpression(`string.length("lixingwang") == 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, true)
}

func TestExpressionGT(t *testing.T) {
	e, err := ParseExpression(`string.length("lixingwang") > 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, false)

	e, err = ParseExpression(`string.length("lixingwang") >= 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, true)

	e, err = ParseExpression(`string.length("lixingwang") < 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, false)

	e, err = ParseExpression(`string.length("lixingwang") <= 10`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	fmt.Println("Result:", v)
	assert.Equal(t, v, true)

}

func TestIsExpression(t *testing.T) {
	b := IsExpression(`string.length("lixingwang") <= 10`)
	assert.True(t, b)

	b = IsExpression(`1>2?string.concat("sss","ddddd"):"fff"`)
	assert.True(t, b)

	b = IsExpression(`string.length("lixingwang")>11?"lixingwang":"fff"`)
	assert.True(t, b)

	b = IsExpression(`string.length("lixingwang")`)
	assert.True(t, b)

	b = IsExpression(`$A3.name.fields`)
	assert.False(t, b)

}

func TestIsTernayExpression(t *testing.T) {
	b := IsExpression(`len("lixingwang") <= 10`)
	assert.True(t, b)

	b = IsExpression(`1>2?concat("sss","ddddd"):"fff"`)
	assert.True(t, b)

	b = IsExpression(`Len("lixingwang")>11?"lixingwang":"fff"`)
	assert.True(t, b)

	b = IsExpression(`len("lixingwang")`)
	assert.True(t, b)

	b = IsExpression(`$A3.name.fields`)
	assert.False(t, b)

}

func TestIsFunction(t *testing.T) {
	b := IsExpression(`len("lixingwang") <= 10`)
	assert.True(t, b)

	b = IsExpression(`1>2?concat("sss","ddddd"):"fff"`)
	assert.True(t, b)

	b = IsExpression(`Len("lixingwang")>11?"lixingwang":"fff"`)
	assert.True(t, b)

	b = IsExpression(`len("lixingwang")`)
	assert.True(t, b)

	b = IsExpression(`$A3.name.fields`)
	assert.False(t, b)
}

func TestNewExpressionBoolean(t *testing.T) {
	e, err := ParseExpression(`(string.length("sea") == 3) == true`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
	fmt.Println("Result:", v)
}

func TestExpressionWithNest(t *testing.T) {
	//Invalid
	e, err := ParseExpression(`(1&&1)==(1&&1)`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	assert.NotNil(t, err)

	//valid case
	e, err = ParseExpression(`(true && true) == false`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, false, v)
}

func TestExpressionWithNILLiteral(t *testing.T) {
	//valid case
	e, err := ParseExpression(`(true && true) != nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`123 != nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`nil == nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
}
func TestExpressionWithNIL(t *testing.T) {
	os.Setenv("name", "test")
	defer func() {
		os.Unsetenv("name")
	}()
	scope := data.NewFixedScope(nil)
	e, err := ParseExpression(`$env.name != nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.EvalWithScope(scope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$env.name == "test"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.EvalWithScope(scope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test == nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", "{}"), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test != nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", `{"test":"123"}`), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test == "123"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", `{"test":"123"}`), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	d := `{"test":"test", "obj":{"id":123, "value":null}}`
	testScope := GetSimpleScope("name", d)

	e, err = ParseExpression(`$.name.test == "test"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.value == nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.id == 123`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.notexist == nil`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
}

func TestExpressionWithNULL(t *testing.T) {
	os.Setenv("name", "test")
	defer func() {
		os.Unsetenv("name")
	}()
	scope := data.NewFixedScope(nil)
	e, err := ParseExpression(`$env.name != null`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.EvalWithScope(scope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$env.name == "test"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.EvalWithScope(scope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test == null`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", "{}"), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test != null`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", `{"test":"123"}`), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.test == "123"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(GetSimpleScope("name", `{"test":"123"}`), data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	d := `{"test":"test", "obj":{"id":123, "value":null}}`
	testScope := GetSimpleScope("name", d)

	e, err = ParseExpression(`$.name.test == "test"`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.value == null`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.id == 123`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)

	e, err = ParseExpression(`$.name.obj.notexist == null`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	v, err = e.EvalWithScope(testScope, data.GetBasicResolver())
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, true, v)
}

func GetSimpleScope(name, value string) data.Scope {
	a, _ := data.NewAttribute(name, data.TypeObject, value)
	maps := make(map[string]*data.Attribute)
	maps[name] = a
	scope := data.NewFixedScope(maps)
	scope.SetAttrValue(name, value)
	return scope
}
