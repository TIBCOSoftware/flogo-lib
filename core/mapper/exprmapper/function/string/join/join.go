package join

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("join-function")

type Join struct {
}

func init() {
	function.Registry(&Join{})
}

func (jo *Join) GetName() string {
	return "join"
}

func (jo *Join) GetCategory() string {
	return "string"
}

func (jo *Join) Eval(a []string, sep string) string {
	// Documentation to function: https://golang.org/pkg/strings/#Join
	log.Debugf("Calling function join with parameters [a = %s, sep = %s]", a, sep)
	return strings.Join(a, sep);
}