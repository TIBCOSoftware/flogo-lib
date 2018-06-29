package contains

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var co = &Contains{}

func TestContains(t *testing.T) {

	sub := co.Eval("Flogo is the most awesome project ever", "awesome")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = co.Eval("Flogo is the most awesome project ever", "XXXXXXXX")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
	
}
