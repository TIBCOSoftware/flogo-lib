package split

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("split-function")

type Split struct {
}

func init() {
	function.Registry(&Split{})
}

func (sp *Split) GetName() string {
	return "split"
}

func (sp *Split) GetCategory() string {
	return "string"
}

func (sp *Split) Eval(s string, sep string) []string {
	// Documentation to function: https://golang.org/pkg/strings/#Split
	log.Debugf("Calling function split with parameters [s = %s, sep = %s]", s, sep)
	return strings.Split(s, sep);
}