package trim

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trim-function")

type Trim struct {
}

func init() {
	function.Registry(&Trim{})
}

func (tr *Trim) GetName() string {
	return "trim"
}

func (tr *Trim) GetCategory() string {
	return "string"
}

func (tr *Trim) Eval(s string, cutset string) string {
	// Documentation to function: https://golang.org/pkg/strings/#Trim
	log.Debugf("Calling function trim with parameters [s = %s, cutset = %s]", s, cutset)
	return strings.Trim(s, cutset);
}