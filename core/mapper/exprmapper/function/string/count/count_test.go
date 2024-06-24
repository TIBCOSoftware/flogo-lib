package count

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var co = &Count{}

func TestCount(t *testing.T) {

	sub := co.Eval("Flogo is the most awesome project ever", "Flogo")
	fmt.Printf("Result [%d] should be equal to: 1\n", sub)
	assert.Equal(t, 1, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "flogo")
	fmt.Printf("Result [%d] should be equal to: 0\n", sub)
	assert.Equal(t, 0, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "XXXXXXXX")
	fmt.Printf("Result [%d] should be equal to: 0\n", sub)
	assert.Equal(t, 0, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "o")
	fmt.Printf("Result [%d] should be equal to: 5\n", sub)
	assert.Equal(t, 5, sub)
}
