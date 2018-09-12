package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"os"
	"github.com/TIBCOSoftware/flogo-lib/config"
)

func TestCoerceToString(t *testing.T) {

	os.Setenv(config.ENV_DATA_SECRET_KEY_KEY, "SecretKey")
	defer func() {
		os.Unsetenv(config.ENV_DATA_SECRET_KEY_KEY)
	}()

	var valInt interface{} = 2
	cval, _ := CoerceToString(valInt)
	assert.Equal(t, "2", cval, "not equal")

	var valBool interface{} = true
	cval, _ = CoerceToString(valBool)
	assert.Equal(t, "true", cval, "not equal")

	var valStr interface{} = "12"
	cval, _ = CoerceToString(valStr)
	assert.Equal(t, "12", cval, "not equal")

	var valFloat interface{} = 1.23
	cval, _ = CoerceToString(valFloat)
	assert.Equal(t, "1.23", cval, "not equal")

	var valNil interface{} // = nil
	cval, _ = CoerceToString(valNil)
	assert.Equal(t, "", cval, "not equal")
}

func TestCoerceToInteger(t *testing.T) {

	var valInt interface{} = 2
	cval, _ := CoerceToInteger(valInt)
	assert.Equal(t, 2, cval, "not equal")

	var valBool interface{} = true
	cval, _ = CoerceToInteger(valBool)
	assert.Equal(t, 1, cval, "not equal")

	var valStr interface{} = "12"
	cval, _ = CoerceToInteger(valStr)
	assert.Equal(t, 12, cval, "not equal")

	var valFloat interface{} = 1.23
	cval, _ = CoerceToInteger(valFloat)
	assert.Equal(t, 1, cval, "not equal")

	var valNil interface{} //= nil
	cval, _ = CoerceToInteger(valNil)
	assert.Equal(t, 0, cval, "not equal")
}

func TestCoerceToNumber(t *testing.T) {

	var valInt interface{} = 2
	cval, _ := CoerceToNumber(valInt)
	assert.Equal(t, 2.0, cval, "not equal")

	var valBool interface{} = true
	cval, _ = CoerceToNumber(valBool)
	assert.Equal(t, 1.0, cval, "not equal")

	var valStr interface{} = "12"
	cval, _ = CoerceToNumber(valStr)
	assert.Equal(t, 12.0, cval, "not equal")

	var valFloat interface{} = 1.23
	cval, _ = CoerceToNumber(valFloat)
	assert.Equal(t, 1.23, cval, "not equal")

	var valNil interface{} //= nil
	cval, _ = CoerceToNumber(valNil)
	assert.Equal(t, 0.0, cval, "not equal")
}

func TestCoerceToBoolean(t *testing.T) {

	var valInt interface{} = 2
	cval, _ := CoerceToBoolean(valInt)
	assert.Equal(t, true, cval, "not equal")

	var valBool interface{} = true
	cval, _ = CoerceToBoolean(valBool)
	assert.Equal(t, true, cval, "not equal")

	var valStr interface{} = "false"
	cval, _ = CoerceToBoolean(valStr)
	assert.Equal(t, false, cval, "not equal")

	var valFloat interface{} = 1.23
	cval, _ = CoerceToBoolean(valFloat)
	assert.Equal(t, true, cval, "not equal")

	var valNil interface{} //= nil
	cval, _ = CoerceToBoolean(valNil)
	assert.Equal(t, false, cval, "not equal")
}

func TestUnmarshComplexObject(t *testing.T) {
	complexStr := `{"metadata":"this is metdata string","value":""}`
	complexObject, err := CoerceToComplexObject(complexStr)
	assert.Nil(t, err)
	assert.NotNil(t, complexObject)
	assert.NotEqual(t, "", complexObject.Value)
}
