package trimprefix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tr = &Trimprefix{}

func TestTrimprefix(t *testing.T) {

	sub := tr.Eval("TIBCO Flogo is the most awesome project EVER...", "TIBCO ")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER...", sub)
	
	sub = tr.Eval("Flogo is the most awesome project EVER...", "TIBCO")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER...", sub)
	
}
