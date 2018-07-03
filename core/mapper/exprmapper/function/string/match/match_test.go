package match

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Match{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval("p([a-z]+)ch", "peach")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.match("p([a-z]+)ch", "peach")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, true, v)
}
