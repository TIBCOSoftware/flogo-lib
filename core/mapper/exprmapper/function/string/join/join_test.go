package join

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jo = &Join{}

func TestJoin(t *testing.T) {

	words := []string{"Flogo","is","the","most","awesome","project","ever"}
	sub := jo.Eval(words, " ")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project ever\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project ever", sub)

	sub = jo.Eval(words, ",")
	fmt.Printf("Result [%s] should be equal to: \"Flogo,is,the,most,awesome,project,ever\"\n", sub)
	assert.Equal(t, "Flogo,is,the,most,awesome,project,ever", sub)
	
}
