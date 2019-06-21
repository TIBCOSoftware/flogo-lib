package title

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ti = &Title{}

func TestTitle(t *testing.T) {

	sub := ti.Eval("Flogo is the most awesome project ever")
	fmt.Printf("Result [%s] should be equal to: \"Flogo Is The Most Awesome Project Ever\"\n", sub)
	assert.Equal(t, "Flogo Is The Most Awesome Project Ever", sub)
	
}
