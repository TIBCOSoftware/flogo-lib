package trim

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tr = &Trim{}

func TestTrim(t *testing.T) {

	sub := tr.Eval("Flogo is the most awesome project EVER...", ".")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER", sub)
	
	sub = tr.Eval("Flogo is the most awesome project EVER...", "F.")
	fmt.Printf("Result [%s] should be equal to: \"logo is the most awesome project EVER\"\n", sub)
	assert.Equal(t, "logo is the most awesome project EVER", sub)
	
	sub = tr.Eval("Flogo is the most awesome project EVER...", "F.l")
	fmt.Printf("Result [%s] should be equal to: \"ogo is the most awesome project EVER\"\n", sub)
	assert.Equal(t, "ogo is the most awesome project EVER", sub)
}
