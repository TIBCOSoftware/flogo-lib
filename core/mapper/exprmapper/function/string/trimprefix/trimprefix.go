package trimprefix

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trimprefix-function")

type Trimprefix struct {
}

func init() {
	function.Registry(&Trimprefix{})
}

func (tr *Trimprefix) GetName() string {
	return "trimprefix"
}

func (tr *Trimprefix) GetCategory() string {
	return "string"
}

func (tr *Trimprefix) Eval(s string, prefix string) string {
	// Documentation to function: https://golang.org/pkg/strings/#TrimPrefix
	log.Debugf("Calling function trimprefix with parameters [s = %s, prefix = %s]", s, prefix)
	return strings.TrimPrefix(s, prefix);
}