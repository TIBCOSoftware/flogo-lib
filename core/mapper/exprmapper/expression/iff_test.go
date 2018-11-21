package expression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIFf(t *testing.T) {
	e, err := ParseExpression(`if 2>1 then "hello"`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "hello", v)
}

func TestIFfExprTrue(t *testing.T) {
	e, err := ParseExpression(`if true then "hello"`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "hello", v)

}

func TestIFfExprStatmentExpr(t *testing.T) {
	e, err := ParseExpression(`if true then 1 > 3`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, false, v)

}

func TestIFfExprStatmentNestIf(t *testing.T) {
	e, err := ParseExpression(`if true then if 2>1 then "yes" `)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "yes", v)
}

func TestIFfElse(t *testing.T) {
	e, err := ParseExpression(`if 1>2 then "yes" else "no"`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

}

func TestIFfElseAllNest(t *testing.T) {
	e, err := ParseExpression(`if (2>1) && (4>3) && (6>5) then if (6>5) && (8 > 7) then "good" else "no" else "no"`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "good", v)

	e, err = ParseExpression(`if (2>1) && (4>3) && (6>5) then if (6>5) && (8 < 7) then "good" else "no" else "no"`)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

	e, err = ParseExpression(`if (2>1) && (4>3) && (6>5) then if (6>5) && (8 < 7) then "good" else "no" else if (6>5) && (8 < 7) then "good" else "no"`)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

	e, err = ParseExpression(`if (2>1) && (4>3) && (6 < 5) then if (6>5) && (8 < 7) then "good" else "no" else if (6>5) && (8 < 7) then "good" else "no"`)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

}
