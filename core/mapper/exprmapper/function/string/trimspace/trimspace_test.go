package trimspace

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tr = &Trimspace{}

func TestTrimspace(t *testing.T) {

	sub := tr.Eval("     TIBCO Flogo is the most awesome project EVER...     ")
	fmt.Printf("Result [%s] should be equal to: \"TIBCO Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "TIBCO Flogo is the most awesome project EVER...", sub)
	
	sub = tr.Eval("TIBCO Flogo is the most awesome project EVER...")
	fmt.Printf("Result [%s] should be equal to: \"TIBCO Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "TIBCO Flogo is the most awesome project EVER...", sub)
}
