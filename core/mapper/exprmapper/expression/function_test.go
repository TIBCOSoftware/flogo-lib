package expression

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/data"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFunctionConcatWithSpace(t *testing.T) {

	e, err := ParseExpression(`flogo.concat("This", "is",string.concat("my","first"),"gocc",string.concat("lexer","and","parser"),string.concat("go","program","!!!"))`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}

	assert.Equal(t, "Thisismyfirstgocclexerandparsergoprogram!!!", function.HandleToSingleOutput(v).(string))
	fmt.Println("Result:", v)
}

func TestFunctionConcatWithMultiSpace(t *testing.T) {

	e, err := ParseExpression(`flogo.concat("This",   " is" , " WI")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "This is WI", function.HandleToSingleOutput(v).(string))
	fmt.Println("Result:", v)
}
func TestFunctionConcat(t *testing.T) {

	e, err := ParseExpression(`flogo.concat("This","is",string.concat("my","first"),"gocc",string.concat("lexer","and","parser"),string.concat("go","program","!!!"))`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "Thisismyfirstgocclexerandparsergoprogram!!!", function.HandleToSingleOutput(v).(string))
	fmt.Println("Result:", v)
}

func TestFunctionLength(t *testing.T) {
	e, err := ParseExpression(`string.length("lixingwang")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, int(10), function.HandleToSingleOutput(v).(int))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionCombine(t *testing.T) {
	e, err := ParseExpression(`string.concat("Beijing",string.tostring(string.length("lixingwang")))`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "Beijing10", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionCombine2(t *testing.T) {
	e, err := ParseExpression(`string.concat("Beijing",string.tostring(string.length("lixingwang")))`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "Beijing10", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionError(t *testing.T) {
	v, err := ParseExpression(`string.concat("Beijing",string.tostring(2017))`)
	if err != nil {
		assert.NotNil(t, err)
		fmt.Println("Result", v)
	} else {
		t.Failed()
	}
}

func TestFunctionWithRefMapping(t *testing.T) {
	e, err := ParseExpression(`string.concat($A3.query.result,"data")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	s, _ := json.Marshal(e)
	fmt.Println(string(s))
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "$A3.query.resultdata", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionWithRefMapping2(t *testing.T) {
	e, err := ParseExpression(`string.concat($A2.message,"lixingwang")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "$A2.messagelixingwang", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionWithTag(t *testing.T) {
	e, err := ParseExpression(`flogo.concat($A2.message,"lixingwang")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "$A2.messagelixingwang", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionWithSpaceInRef(t *testing.T) {
	e, err := ParseExpression(`string.concat($Marketo Get Lead by Id.output.result[0].firstName,$Marketo Get Lead by Id.output.result[0].lastName)`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "$Marketo Get Lead by Id.output.result[0].firstName$Marketo Get Lead by Id.output.result[0].lastName", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionWithPackage(t *testing.T) {
	e, err := ParseExpression(`string.concat($A2.message,"lixingwang")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "$A2.messagelixingwang", function.HandleToSingleOutput(v).(string))

	fmt.Println("Result:", function.HandleToSingleOutput(v))
}

func TestFunctionWithNil(t *testing.T) {
	e, err := ParseExpression(`string.concat("test","nil")`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "testnil", function.HandleToSingleOutput(v).(string))

	e, err = ParseExpression(`string.concat("test",nil)`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err = e.Eval()
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	assert.Equal(t, "test", function.HandleToSingleOutput(v).(string))

}

func TestFunctionWithSpecialFiled(t *testing.T) {
	e, err := ParseExpression(`$activity[lixingwang].myattri["name.name"][0]>2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	_, err = e.Eval()
	assert.NotNil(t, err)
}

func TestFunctionWithNegtiveNumber(t *testing.T) {
	d := `{"test":"test", "obj":{"id":-123, "value":"value"}}`
	testScope := GetSimpleScope("name", d)

	e, err := ParseExpression(`string.length($.name.obj.value) >= -2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	v, err := e.EvalWithScope(testScope, data.GetBasicResolver())
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

type Concat struct {
}

func init() {
	function.Registry(&Concat{})
}

func (s *Concat) GetName() string {
	return "concat"
}

func (s *Concat) GetCategory() string {
	return "flogo"
}

func (s *Concat) Eval(strs ...string) string {
	log.Infof("Start flogo:concat function with parameters %s", strs)
	var buffer bytes.Buffer

	for _, v := range strs {
		buffer.WriteString(v)
	}
	log.Debugf("Done flogo:concat function with result %s", buffer.String())
	return buffer.String()
}

type PConcat struct {
}

func init() {
	function.Registry(&PConcat{})
}

func (s *PConcat) GetName() string {
	return "string.concat"
}

func (s *PConcat) GetCategory() string {
	return ""
}

func (s *PConcat) Eval(strs ...string) string {
	log.Debugf("Start string:concat function with parameters %s", strs)
	var buffer bytes.Buffer

	for _, v := range strs {
		buffer.WriteString(v)
	}
	log.Debugf("Done string:concat function with result %s", buffer.String())
	return buffer.String()
}

type Length struct {
}

func init() {
	function.Registry(&Length{})
}

func (s *Length) GetName() string {
	return "length"
}

func (s *Length) GetCategory() string {
	return "string"
}

func (s *Length) Eval(str string) int {
	log.Debugf("Return the length of a string \"%s\"", str)
	var l int
	//l = len([]rune(str))
	l = len(str)
	log.Debugf("Done calculating the length %d", l)
	return l
}

type PanicFunc struct {
}

func init() {
	function.Registry(&PanicFunc{})
}

func (s *PanicFunc) GetName() string {
	return "panic"
}

func (s *PanicFunc) GetCategory() string {
	return "panic"
}

func (s *PanicFunc) Eval() string {
	panic("Panic happened")
	return "panic"
}

func TestPanictFunction(t *testing.T) {
	e, err := ParseExpression(`panic.panic()`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.NotNil(t, err)
	assert.Nil(t, v)
}

func TestNumberLenFunction(t *testing.T) {
	v, err := ParseExpression(`string.length("hello,world")`)
	log.Info(v)
	assert.NotNil(t, v)
	assert.Nil(t, err)

}

type String struct {
}

func init() {
	function.Registry(&String{})
}

func (s *String) GetName() string {
	return "tostring"
}

func (s *String) GetCategory() string {
	return "string"
}

func (s *String) Eval(in interface{}) string {
	log.Debugf("Start String function with parameters %s", in)

	switch in.(type) {
	case string:
		return in.(string)
	case float64:
		return strconv.FormatFloat(in.(float64), 'f', -1, 64)
	case int32, int8, int16, int64:
		//v := int64(in)
		return strconv.FormatInt(in.(int64), 10)
	case int:
		return strconv.Itoa(in.(int))
	case *int:
		return strconv.Itoa(*in.(*int))
	case uint, uint8, uint16, uint32, uint64:
		//v := int64(in)
		return strconv.FormatInt(in.(int64), 10)
	default:
		str, err := data.CoerceToString(in)
		if err != nil {
			log.Errorf("Convert to string error %s", err.Error())
		}
		return str
	}
}
