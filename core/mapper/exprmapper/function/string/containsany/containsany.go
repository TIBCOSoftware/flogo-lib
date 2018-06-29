package containsany

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("containsany-function")

type Containsany struct {
}

func init() {
	function.Registry(&Containsany{})
}

func (co *Containsany) GetName() string {
	return "containsany"
}

func (co *Containsany) GetCategory() string {
	return "string"
}

func (co *Containsany) Eval(s string, chars string) bool {
	// Documentation to function: https://golang.org/pkg/strings/#ContainsAny
	log.Debugf("Calling function containsany with parameters [s = %s, chars = %s]", s, chars)
	return strings.ContainsAny(s, chars);
}