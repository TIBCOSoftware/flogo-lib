package trimsuffix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tr = &Trimsuffix{}

func TestTrimsuffix(t *testing.T) {

	sub := tr.Eval("TIBCO Flogo is the most awesome project EVER...", "TIBCO ")
	fmt.Printf("Result [%s] should be equal to: \"TIBCO Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "TIBCO Flogo is the most awesome project EVER...", sub)
	
	sub = tr.Eval("Flogo is the most awesome project EVER...", " EVER...")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project", sub)
	
	sub = tr.Eval("Flogo is the most awesome project EVER...", "")
	fmt.Printf("Result [%s] should be equal to: \"Flogo is the most awesome project EVER...\"\n", sub)
	assert.Equal(t, "Flogo is the most awesome project EVER...", sub)
}
