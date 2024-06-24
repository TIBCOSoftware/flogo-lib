package repeat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var re = &Repeat{}

func TestRepeat(t *testing.T) {

	sub := "Flogo" + re.Eval("go", 5)
	fmt.Printf("Result [%s] should be equal to: \"Flogogogogogogo\"\n", sub)
	assert.Equal(t, "Flogogogogogogo", sub)
	
	sub = "Flogo is the most awesome project" + re.Eval(" ever", 2)
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project ever ever\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project ever ever", sub)
}
