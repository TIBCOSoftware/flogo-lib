package expr

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/ref"
	"reflect"
)

type LiteralExpr struct {
	V interface{}
}

func NewLiteralExpr(v interface{}) *LiteralExpr {
	return &LiteralExpr{V: v}
}
func (iffl *LiteralExpr) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	return iffl.EvalWithData(nil, inputScope, resolver)
}
func (iffl *LiteralExpr) Eval() (interface{}, error) {
	return iffl.EvalWithData(nil, nil, nil)
}
func (iffl *LiteralExpr) EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	switch t := iffl.V.(type) {
	case *ref.ArrayRef:
		if inputScope == nil {
			return t.GetRef(), nil
		} else {
			if value == nil {
				//Array mapping should not go here for today, take is as get current scope.
				//TODO how to know it is array mapping or get current scope
				ref := ref.NewMappingRef(t.GetRef())
				v, err := ref.Eval(inputScope, resolver)
				if err != nil {
					return nil, err
				}
				return v, nil
			} else {
				v, err := t.EvalFromData(value)
				if err != nil {
					return reflect.Value{}, err
				}
				return v, nil
			}

		}

		return handleArrayRef(value, t.GetRef(), inputScope, resolver)
	case *ref.MappingRef:
		if inputScope == nil {
			return t.GetRef(), nil
		} else {

			v, err := t.Eval(inputScope, resolver)
			if err != nil {
				return reflect.Value{}, err
			}
			return v, nil
		}
	default:
		return iffl.V, nil
	}
}
