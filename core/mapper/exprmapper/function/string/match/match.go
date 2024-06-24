package match

import (
	"regexp"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("match-function")

type Match struct {
}

func init() {
	function.Registry(&Match{})
}

func (s *Match) GetName() string {
	return "match"
}

func (s *Match) GetCategory() string {
	return "string"
}
func (s *Match) Eval(expr string, inp string) bool {

	match, _ := regexp.MatchString(expr, inp)
	log.Infof(`Input "%v" matches: "%v" `, inp, match)
	return match
}
