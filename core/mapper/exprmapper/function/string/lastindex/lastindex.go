package lastindex

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("lastindex-function")

type Lastindex struct {
}

func init() {
	function.Registry(&Lastindex{})
}

func (la *Lastindex) GetName() string {
	return "lastindex"
}

func (la *Lastindex) GetCategory() string {
	return "string"
}

func (la *Lastindex) Eval(s string, substr string) int {
	// Documentation to function: https://golang.org/pkg/strings/#LastIndex
	log.Debugf("Calling function lastindex with parameters [s = %s, substr = %s]", s, substr)
	return strings.LastIndex(s, substr);
}