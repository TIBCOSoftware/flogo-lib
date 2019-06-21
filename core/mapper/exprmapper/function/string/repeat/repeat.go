package repeat

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("repeat-function")

type Repeat struct {
}

func init() {
	function.Registry(&Repeat{})
}

func (re *Repeat) GetName() string {
	return "repeat"
}

func (re *Repeat) GetCategory() string {
	return "string"
}

func (re *Repeat) Eval(s string, count int) string {
	// Documentation to function: https://golang.org/pkg/strings/#Repeat
	log.Debugf("Calling function repeat with parameters [s = %s, count = %s]", s, count)
	return strings.Repeat(s, count);
}