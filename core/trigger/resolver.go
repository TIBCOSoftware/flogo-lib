package trigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func init() {
	data.SetResolver(data.RES_TRIGGER, Resolve)
}

// Resolve will resolve a trigger output value in the given scope
func Resolve(scope data.Scope, value string) (interface{}, bool) {
	attr, ok := scope.GetAttr("{T." + value + "}")
	if !ok {
		return nil, false
	}

	return attr.Value, true
}
