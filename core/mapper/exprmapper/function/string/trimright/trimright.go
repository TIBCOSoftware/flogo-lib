package trimright

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("trimright-function")

type Trimright struct {
}

func init() {
	function.Registry(&Trimright{})
}

func (tr *Trimright) GetName() string {
	return "trimright"
}

func (tr *Trimright) GetCategory() string {
	return "string"
}

func (tr *Trimright) Eval(s string, cutset string) string {
	// Documentation to function: https://golang.org/pkg/strings/#TrimRight
	log.Debugf("Calling function trimright with parameters [s = %s, cutset = %s]", s, cutset)
	return strings.TrimRight(s, cutset);
}