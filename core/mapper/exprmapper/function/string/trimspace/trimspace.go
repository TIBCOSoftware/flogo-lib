package trimspace

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trimspace-function")

type Trimspace struct {
}

func init() {
	function.Registry(&Trimspace{})
}

func (tr *Trimspace) GetName() string {
	return "trimspace"
}

func (tr *Trimspace) GetCategory() string {
	return "string"
}

func (tr *Trimspace) Eval(s string) string {
	// Documentation to function: https://golang.org/pkg/strings/#TrimSpace
	log.Debugf("Calling function trimspace with parameters [s = %s]", s)
	return strings.TrimSpace(s);
}