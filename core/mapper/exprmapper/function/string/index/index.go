package index

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)

var log = logger.GetLogger("index-function")

type Index struct {
}

func init() {
	function.Registry(&Index{})
}

func (in *Index) GetName() string {
	return "index"
}

func (in *Index) GetCategory() string {
	return "string"
}

func (in *Index) Eval(s string, substr string) int {
	// Documentation to function: https://golang.org/pkg/strings/#Index
	log.Debugf("Calling function index with parameters [s = %s, substr = %s]", s, substr)
	return strings.Index(s, substr);
}