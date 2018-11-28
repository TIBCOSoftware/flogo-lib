package expr

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/ref"
)

type IfValueProvider struct {
	V interface{}
}

func NewIfValueProvider(v interface{}) *IfValueProvider {
	return &IfValueProvider{V: v}
}

func (iffl *IfValueProvider) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	return iffl.V, nil
}

func (iffl *IfValueProvider) Eval() (interface{}, error) {
	return iffl.V, nil
}

func (iffl *IfValueProvider) EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	switch t := iffl.V.(type) {

	case *ref.ArrayRef:
		return handleArrayRef(value, t.GetRef(), inputScope, resolver)
	case *ref.MappingRef:
		return t.Eval(inputScope, resolver)
	default:
		return iffl.V, nil
	}
}
