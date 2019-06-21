package hasprefix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ha = &Hasprefix{}

func TestHasprefix(t *testing.T) {

	sub := ha.Eval("Flogo is the most awesome project ever", "Flogo")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "flogo")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "Any other thing")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
}
