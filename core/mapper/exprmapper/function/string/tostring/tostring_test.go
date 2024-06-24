package tostring

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &ToString{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval(1)
	fmt.Println(final1)
	assert.Equal(t, "1", final1)

	final2 := s.Eval(1.23)
	fmt.Println(final2)
	assert.Equal(t, "1.23", final2)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.tostring(3.456)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, "3.456", v)
}
