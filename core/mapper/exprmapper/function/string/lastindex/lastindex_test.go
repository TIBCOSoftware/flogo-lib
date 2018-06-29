package lastindex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var la = &Lastindex{}

func TestLastindex(t *testing.T) {

	sub := la.Eval("Flogo is the most awesome project ever", "o")
	fmt.Printf("Result [%d] should be equal to: 28\n", sub)
	assert.Equal(t, 28, sub)
	
	sub = la.Eval("Flogo is the most awesome project ever", "x")
	fmt.Printf("Result [%d] should be equal to: -1\n", sub)
	assert.Equal(t, -1, sub)
	
	sub = la.Eval("Flogo is the most awesome project ever", "flogo")
	fmt.Printf("Result [%d] should be equal to: -1\n", sub)
	assert.Equal(t, -1, sub)
	
	sub = la.Eval("Flogo is the most awesome project ever", "Flogo")
	fmt.Printf("Result [%d] should be equal to: 0\n", sub)
	assert.Equal(t, 0, sub)
}
