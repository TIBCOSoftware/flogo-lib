package replace

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("replace-function")

type Replace struct {
}

func init() {
	function.Registry(&Replace{})
}

func (re *Replace) GetName() string {
	return "replace"
}

func (re *Replace) GetCategory() string {
	return "string"
}

func (re *Replace) Eval(s string, old string, new string, n int) string {
	// Documentation to function: https://golang.org/pkg/strings/#Replace
	log.Debugf("Calling function replace with parameters [s = %s, old = %s, new = %s, n = %s]", s, old, new, n)
	return strings.Replace(s, old, new, n);
}