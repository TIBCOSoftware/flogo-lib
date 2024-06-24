package toupper

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("toupper-function")

type Toupper struct {
}

func init() {
	function.Registry(&Toupper{})
}

func (to *Toupper) GetName() string {
	return "toupper"
}

func (to *Toupper) GetCategory() string {
	return "string"
}

func (to *Toupper) Eval(s string) string {
	// Documentation to function: https://golang.org/pkg/strings/#ToUpper
	log.Debugf("Calling function toupper with parameters [s = %s]", s)
	return strings.ToUpper(s);
}