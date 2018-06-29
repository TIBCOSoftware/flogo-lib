package contains

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("contains-function")

type Contains struct {
}

func init() {
	function.Registry(&Contains{})
}

func (co *Contains) GetName() string {
	return "contains"
}

func (co *Contains) GetCategory() string {
	return "string"
}

func (co *Contains) Eval(s string, substr string) bool {
	// Documentation to function: https://golang.org/pkg/strings/#Contains
	log.Debugf("Calling function contains with parameters [s = %s, substr = %s]", s, substr)
	return strings.Contains(s, substr);
}