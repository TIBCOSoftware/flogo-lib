package trimleft

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trimleft-function")

type Trimleft struct {
}

func init() {
	function.Registry(&Trimleft{})
}

func (tr *Trimleft) GetName() string {
	return "trimleft"
}

func (tr *Trimleft) GetCategory() string {
	return "string"
}

func (tr *Trimleft) Eval(s string, cutset string) string {
	// Documentation to function: https://golang.org/pkg/strings/#TrimLeft
	log.Debugf("Calling function trimleft with parameters [s = %s, cutset = %s]", s, cutset)
	return strings.TrimLeft(s, cutset);
}