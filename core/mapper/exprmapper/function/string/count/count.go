package count

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("count-function")

type Count struct {
}

func init() {
	function.Registry(&Count{})
}

func (co *Count) GetName() string {
	return "count"
}

func (co *Count) GetCategory() string {
	return "string"
}

func (co *Count) Eval(s string, substr string) int {
	// Documentation to function: https://golang.org/pkg/strings/#Count
	log.Debugf("Calling function count with parameters [s = %s, substr = %s]", s, substr)
	return strings.Count(s, substr);
}