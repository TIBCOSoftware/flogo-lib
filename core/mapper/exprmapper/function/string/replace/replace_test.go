package replace

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var re = &Replace{}

func TestReplace(t *testing.T) {

	sub := re.Eval("Flogo is the most awesome project ever", "ever", "EVER", -1)
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER", sub)
	
	sub = re.Eval("Flogo is the most awesome project ever ever ever", "ever", "EVER", 2)
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER EVER ever\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER EVER ever", sub)
	
	sub = re.Eval("Flogo is the most awesome project ever", " ever", ". PERIOD", 0)
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project ever\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project ever", sub)
	
}
