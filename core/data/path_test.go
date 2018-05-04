package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathGetValue(t *testing.T) {
	// Resolution of Old Trigger expression

	mapVal, _ := CoerceToObject("{\"myParam\":5}")
	path := ".myParam"
	newVal, err := PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, newVal)

	// Resolution of Old Trigger expression
	arrVal, _ := CoerceToArray("[1,6,3]")
	path = "[1]"
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 6.0, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = ".myParam.nestedMap"
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = `["myParam"].nestedMap`
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = `.myParam["nestedMap"]`
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, newVal)

	arrVal, _ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":2}]")
	path = "[1].nestedMap2"
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 2.0, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedArray\":[7,8,9]}}")
	path = ".myParam.nestedArray[1]"
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 8.0, newVal)

	arrVal, _ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":{\"nestedArray\":[7,8,9]}}]")
	path = "[1].nestedMap2.nestedArray[2]"
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 9.0, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedArray\":[7,8,9]}}")
	path = ".myParam.nestedArray"
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	//todo check if array

	arrVal, _ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":{\"nestedArray\":[7,8,9]}}]")
	path = "[1].nestedMap2"
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	//todo check if map

}

func TestPathSetValue(t *testing.T) {
	// Resolution of Old Trigger expression

	mapVal, _ := CoerceToObject("{\"myParam\":5}")
	path := ".myParam"
	err := PathSetValue(mapVal, path, 6)
	assert.Nil(t, err)
	newVal, err := PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 6, newVal)

	// Resolution of Old Trigger expression
	arrVal, _ := CoerceToArray("[1,6,3]")
	path = "[1]"
	err = PathSetValue(arrVal, path, 4)
	assert.Nil(t, err)
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 4, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = ".myParam.nestedMap"
	assert.Nil(t, err)
	err = PathSetValue(mapVal, path, 7)
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 7, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = `["myParam"].nestedMap`
	assert.Nil(t, err)
	err = PathSetValue(mapVal, path, 7)
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 7, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedMap\":1}}")
	path = `.myParam["nestedMap"]`
	assert.Nil(t, err)
	err = PathSetValue(mapVal, path, 7)
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 7, newVal)


	arrVal, _ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":2}]")
	path = "[1].nestedMap2"
	err = PathSetValue(arrVal, path, 3)
	assert.Nil(t, err)
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 3, newVal)

	mapVal, _ = CoerceToObject("{\"myParam\":{\"nestedArray\":[7,8,9]}}")
	path = ".myParam.nestedArray[1]"
	err = PathSetValue(mapVal, path, 1)
	assert.Nil(t, err)
	newVal, err = PathGetValue(mapVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 1, newVal)

	arrVal, _ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":{\"nestedArray\":[7,8,9]}}]")
	path = "[1].nestedMap2.nestedArray[2]"
	err = PathSetValue(arrVal, path, 5)
	assert.Nil(t, err)
	newVal, err = PathGetValue(arrVal, path)
	assert.Nil(t, err)
	assert.Equal(t, 5, newVal)

	//mapVal,_ = CoerceToObject("{\"myParam\":{\"nestedArray\":[7,8,9]}}")
	//path = ".myParam.nestedArray"
	//err = PathSetValue(arrVal, path, 3)
	//assert.Nil(t, err)
	//newVal,err = PathGetValue(mapVal, path)
	//assert.Nil(t, err)
	////todo check if array
	//
	//arrVal,_ = CoerceToArray("[{\"nestedMap1\":1},{\"nestedMap2\":{\"nestedArray\":[7,8,9]}}]")
	//path = "[1].nestedMap2"
	//assert.Nil(t, err)
	//err = PathSetValue(arrVal, path, 3)
	//newVal,err = PathGetValue(arrVal, path)
	//assert.Nil(t, err)
	//////todo check if map
}
