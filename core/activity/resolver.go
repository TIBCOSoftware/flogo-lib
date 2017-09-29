package activity

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func init() {
	data.SetResolver(data.RES_ACTIVITY, Resolve)
}

// Resolve will resolve a activity output value in the given scope
func Resolve(scope data.Scope, value string) (interface{}, bool) {
	attr, ok := scope.GetAttr("{A." + value + "}")
	if !ok {
		return nil, false
	}

	return attr.Value, true
}

//todo add a "add to scope"