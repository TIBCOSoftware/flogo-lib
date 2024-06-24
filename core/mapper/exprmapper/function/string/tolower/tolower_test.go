package tolower

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var to = &Tolower{}

func TestTolower(t *testing.T) {

	sub := to.Eval("Flogo is the most awesome project EVER")
	fmt.Printf("Result [%s] should be equal to: \"flogo is the most awesome project ever\"\n", sub)
	assert.Equal(t, "flogo is the most awesome project ever", sub)
}
