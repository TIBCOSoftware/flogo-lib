package hasprefix

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("hasprefix-function")

type Hasprefix struct {
}

func init() {
	function.Registry(&Hasprefix{})
}

func (ha *Hasprefix) GetName() string {
	return "hasprefix"
}

func (ha *Hasprefix) GetCategory() string {
	return "string"
}

func (ha *Hasprefix) Eval(s string, prefix string) bool {
	// Documentation to function: https://golang.org/pkg/strings/#HasPrefix
	log.Debugf("Calling function hasprefix with parameters [s = %s, prefix = %s]", s, prefix)
	return strings.HasPrefix(s, prefix);
}