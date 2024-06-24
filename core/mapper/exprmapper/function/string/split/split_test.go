package split

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sp = &Split{}

func TestSplit(t *testing.T) {

	sub := sp.Eval("Flogo is the most awesome project ever", " ")
	chk := []string{"Flogo","is","the","most","awesome","project","ever"}
	fmt.Printf("Result [%s] should be equal to: \"[Flogo is the most awesome project ever]\" and size should be [7]\n", sub)
	
	assert.Equal(t, chk, sub)
	assert.Equal(t, len(sub), 7)
		
	sub = sp.Eval("Flogo is the most awesome project ever", "-")
	chk = []string{"Flogo is the most awesome project ever"}
	fmt.Printf("Result [%s] should be equal to: \"[Flogo is the most awesome project ever]\" and size should be [1]\"\n", sub)
	
	assert.Equal(t, chk, sub)
	assert.Equal(t, len(sub), 1)
	
	
}
