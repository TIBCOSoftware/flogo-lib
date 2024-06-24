package title

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("title-function")

type Title struct {
}

func init() {
	function.Registry(&Title{})
}

func (ti *Title) GetName() string {
	return "title"
}

func (ti *Title) GetCategory() string {
	return "string"
}

func (ti *Title) Eval(s string) string {
	// Documentation to function: https://golang.org/pkg/strings/#Title
	log.Debugf("Calling function title with parameters [s = %s]", s)
	return strings.Title(s);
}