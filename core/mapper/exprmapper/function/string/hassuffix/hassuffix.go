package hassuffix

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("hassuffix-function")

type Hassuffix struct {
}

func init() {
	function.Registry(&Hassuffix{})
}

func (ha *Hassuffix) GetName() string {
	return "hassuffix"
}

func (ha *Hassuffix) GetCategory() string {
	return "string"
}

func (ha *Hassuffix) Eval(s string, suffix string) bool {
	// Documentation to function: https://golang.org/pkg/strings/#HasSuffix
	log.Debugf("Calling function hassuffix with parameters [s = %s, suffix = %s]", s, suffix)
	return strings.HasSuffix(s, suffix);
}