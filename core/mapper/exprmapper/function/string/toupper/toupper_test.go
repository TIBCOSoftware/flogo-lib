package toupper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var to = &Toupper{}

func TestToupper(t *testing.T) {

	sub := to.Eval("Flogo is the most awesome project EVER")
	fmt.Printf("Result [%s] should be equal to: \"FLOGO IS THE MOST AWESOME PROJECT EVER\"\n", sub)
	assert.Equal(t, "FLOGO IS THE MOST AWESOME PROJECT EVER", sub)
}
