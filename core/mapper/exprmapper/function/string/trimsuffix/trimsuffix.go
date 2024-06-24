package trimsuffix

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trimsuffix-function")

type Trimsuffix struct {
}

func init() {
	function.Registry(&Trimsuffix{})
}

func (tr *Trimsuffix) GetName() string {
	return "trimsuffix"
}

func (tr *Trimsuffix) GetCategory() string {
	return "string"
}

func (tr *Trimsuffix) Eval(s string, suffix string) string {
	// Documentation to function: https://golang.org/pkg/strings/#TrimSuffix
	log.Debugf("Calling function trimsuffix with parameters [s = %s, suffix = %s]", s, suffix)
	return strings.TrimSuffix(s, suffix);
}