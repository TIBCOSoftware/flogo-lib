package tolower

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("tolower-function")

type Tolower struct {
}

func init() {
	function.Registry(&Tolower{})
}

func (to *Tolower) GetName() string {
	return "tolower"
}

func (to *Tolower) GetCategory() string {
	return "string"
}

func (to *Tolower) Eval(s string) string {
	// Documentation to function: https://golang.org/pkg/strings/#ToLower
	log.Debugf("Calling function tolower with parameters [s = %s]", s)
	return strings.ToLower(s);
}