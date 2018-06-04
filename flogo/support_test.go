package flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMappingValue(t *testing.T) {

	//todo add additional tests when support for more mapping type is added
	mType, mValue := getMappingValue("$.value")

	assert.Equal(t, data.MtExpression, mType)
	assert.Equal(t, "$.value", mValue)
}

func TestToMappings(t *testing.T) {

	mappings := []string{"in1=b", "in2= $.blah", "in3 = $.blah2"}

	//todo add additional tests when support for more mapping type is added
	defs, err := toMappingDefs(mappings)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(defs))

	assert.Equal(t, "in1", defs[0].MapTo)
	assert.Equal(t, "in2", defs[1].MapTo)
	assert.Equal(t, "in3", defs[2].MapTo)

	assert.Equal(t, data.MtExpression, defs[0].Type)
	assert.Equal(t, data.MtExpression, defs[1].Type)
	assert.Equal(t, data.MtExpression, defs[2].Type)

	assert.Equal(t, "b", defs[0].Value)
	assert.Equal(t, "$.blah", defs[1].Value)
	assert.Equal(t, "$.blah2", defs[2].Value)

}
