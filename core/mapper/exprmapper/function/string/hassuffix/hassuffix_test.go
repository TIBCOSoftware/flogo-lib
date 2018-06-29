package hassuffix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ha = &Hassuffix{}

func TestHassuffix(t *testing.T) {

	sub := ha.Eval("Flogo is the most awesome project ever", "ever")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "EVER")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "XXXX")
	fmt.Printf("Result [%t] should be equal to: false\n", sub)
	assert.Equal(t, false, sub)
	
	sub = ha.Eval("Flogo is the most awesome project ever", "")
	fmt.Printf("Result [%t] should be equal to: true\n", sub)
	assert.Equal(t, true, sub)

}