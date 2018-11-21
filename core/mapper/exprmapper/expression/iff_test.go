package expression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIFf(t *testing.T) {
	ifExpr := `
if 2>1 {
	"hello"
}`
	e, err := ParseExpression(ifExpr)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "hello", v)
}

func TestIFfExprTrue(t *testing.T) {
	e, err := ParseExpression(`if true { "hello" }`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "hello", v)

}

func TestIFfExprStatmentExpr(t *testing.T) {
	e, err := ParseExpression(`if true { 1 > 3 }`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, false, v)

}

func TestIFfExprStatmentNestIf(t *testing.T) {
	e, err := ParseExpression(`if true { if 2>1 { "yes" } }`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "yes", v)
}

func TestIFfElse(t *testing.T) {
	e, err := ParseExpression(`if 1>2 { "yes" } else { "no" }`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

}

func TestIFfElseAllNest(t *testing.T) {
	e, err := ParseExpression(`if (2>1) && (4>3) && (6>5) { if (6>5) && (8 > 7) { "good" } else  { "no" } } else { "no" }`)
	assert.Nil(t, err)
	v, err := e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "good", v)

	e, err = ParseExpression(`if (2>1) && (4>3) && (6>5) {  if (6>5) && (8 < 7) { "good" } else {"no" } } else {"no"}`)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

	e, err = ParseExpression(`if (2>1) && (4>3) && (6>5) { if (6>5) && (8 < 7) { "good" } else { "no"} } else { if (6>5) && (8 < 7) { "good" } else { "no" } }`)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

	ifExpr := `
if (2>1) && (4>3) && (6 < 5) { 
	if (6>5) && (8 < 7) {
		 "good"
    } else {
    	 "no"
    }
} else {
	if (6>5) && (8 < 7) { 
		"good"
	} else {
		"no"
	}
}
`
	e, err = ParseExpression(ifExpr)
	assert.Nil(t, err)
	v, err = e.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "no", v)

}
