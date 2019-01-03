package expression

import (
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/number/Int64"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperatorPrecedenceForMulAdd(t *testing.T) {

	v, err := ParseExpression(`1 + 2 * 3 + 2 * 6 / 2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, err := v.EvalWithScope(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 13, vv)

	v, err = ParseExpression(` 1 + 4 * 5 + -6 `)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, err = v.EvalWithScope(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 15, vv)

	v, err = ParseExpression(` 2 < 3 && 5 > 4 && 6 < 7 && 56 > 44`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, err = v.EvalWithScope(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, true, vv)

	v, err = ParseExpression(` 2 < 3 && 5 > 4 ||  6 < 7 && 56 < 44`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, err = v.EvalWithScope(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, true, vv)

}

func TestOperatorMod(t *testing.T) {
	//number.int64($flow.pathParams.id)-2 >=1
	v, err := ParseExpression(`3-2`)
	if err != nil {
		t.Fatal(err)
		t.Failed()
	}
	vv, err := v.EvalWithScope(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, vv)
}
