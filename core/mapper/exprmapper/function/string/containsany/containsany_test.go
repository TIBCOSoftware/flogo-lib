package containsany

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var co = &Containsany{}

func TestContainsany(t *testing.T) {

	sub := co.Eval("Flogo is the most awesome project ever", "o")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "i & z")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "z & go")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "z")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
}
