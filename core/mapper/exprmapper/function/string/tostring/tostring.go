package tostring

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("equals-function")

type ToString struct {
}

func init() {
	function.Registry(&ToString{})
}

func (s *ToString) GetName() string {
	return "toString"
}

func (s *ToString) GetCategory() string {
	return "string"
}
func (s *ToString) Eval(inp interface{}) string {

	result := fmt.Sprintf("%v", inp)
	log.Debugf(`Converts value to a simple string: "%s" `, result)
	return result
}
