package index

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var in = &Index{}

func TestIndex(t *testing.T) {

	sub := in.Eval("Flogo is the most awesome project ever", "Flogo")
	fmt.Printf("Result [%d] should be equal to: 0\n", sub)
	assert.Equal(t, 0, sub)

	sub = in.Eval("Flogo is the most awesome project ever", "flogo")
	fmt.Printf("Result [%d] should be equal to: -1\n", sub)
	assert.Equal(t, -1, sub)
	
	sub = in.Eval("Flogo is the most awesome project ever", "awesome")
	fmt.Printf("Result [%d] should be equal to: 18\n", sub)
	assert.Equal(t, 18, sub)
	
}
