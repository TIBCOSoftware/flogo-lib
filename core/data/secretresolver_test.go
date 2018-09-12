package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"os"
	"github.com/TIBCOSoftware/flogo-lib/config"
)

func TestSecretKeyDefault(t *testing.T) {
	defer func() {
		SetSecretValueHandler(nil)
	}()
	handler := GetSecretValueHandler()
	encoded, err := handler.EncodeValue("mysecurepassword3")
	assert.Nil(t, err)
	decoded, err := handler.DecodeValue(encoded)
	assert.Nil(t, err)
	assert.Equal(t, "mysecurepassword3", decoded)
}

func TestSecretKeyEnv(t *testing.T) {
	os.Setenv(config.ENV_DATA_SECRET_KEY_KEY, "mysecretkey1")
	defer func() {
		os.Unsetenv(config.ENV_DATA_SECRET_KEY_KEY)
		SetSecretValueHandler(nil)
	}()

	handler := GetSecretValueHandler()
	encoded, err := handler.EncodeValue("mysecurepassword1")
	assert.Nil(t, err)
	decoded, err := handler.DecodeValue(encoded)
	assert.Nil(t, err)
	assert.Equal(t, "mysecurepassword1", decoded)
}

func TestSecretKey(t *testing.T) {
	defer func() {
		SetSecretValueHandler(nil)
	}()
	SetSecretValueHandler(&KeyBasedSecretValueHandler{Key: "mysecretkey2"})
	handler := GetSecretValueHandler()
	encoded, err := handler.EncodeValue("mysecurepassword1")
	assert.Nil(t, err)
	decoded, err := GetSecretValueHandler().DecodeValue(encoded)
	assert.Nil(t, err)
	assert.Equal(t, "mysecurepassword1", decoded)
}




