package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCmdOverrideAppProperty(t *testing.T) {

	//value with quotes and ,
	m := GetCmdOverrideAppProperty(`name="select id, name , ssss",val=ddd`)
	assert.Equal(t, "select id, name , ssss", m["name"])
	assert.Equal(t, "ddd", m["val"])

	//value with =
	m = GetCmdOverrideAppProperty(`name="dddd",val="ddd"`)
	assert.Equal(t, "dddd", m["name"])
	assert.Equal(t, "ddd", m["val"])

	m = GetCmdOverrideAppProperty(`name="name=v",val="select=v"`)

	assert.Equal(t, "name=v", m["name"])
	assert.Equal(t, "select=v", m["val"])

	m = GetCmdOverrideAppProperty(`name=select name=name,val=select`)

	assert.Equal(t, "select name=name", m["name"])
	assert.Equal(t, "select", m["val"])

	m = GetCmdOverrideAppProperty(`name=select name,name,val=select=n`)

	assert.NotEqual(t, "select name=name", m["name"])
	assert.NotEqual(t, "select=n", m["val"])

	//Space
	m = GetCmdOverrideAppProperty(`na me = select name=name,va l = select`)

	assert.Equal(t, "select name=name", m["na me"])
	assert.Equal(t, "select", m["va l"])

	//Empty value
	m = GetCmdOverrideAppProperty(`MyPassword=SECRET:sjsjsjsjsjsarebbxmooensgajkakap/hsjsjsjks==,MyVal=test`)

	assert.Equal(t, "SECRET:sjsjsjsjsjsarebbxmooensgajkakap/hsjsjsjks==", m["MyPassword"])
	assert.Equal(t, "test", m["MyVal"])

}
