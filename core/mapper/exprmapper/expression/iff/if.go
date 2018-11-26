package iff

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/expr"
)

type IFF interface {
	EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error)
	Eval() (interface{}, error)
	EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error)
}

type iff struct {
	Expr expr.Expr
	Then expr.Expr
	Else expr.Expr
}

func NewIff(expr, then expr.Expr) *iff {
	return &iff{Expr: expr, Then: then}
}

func NewIffElse(expr, then, elseS expr.Expr) *iff {
	return &iff{Expr: expr, Then: then, Else: elseS}
}

func (iff *iff) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	exprValue, err := iff.Expr.EvalWithScope(inputScope, resolver)
	if err != nil {
		return nil, fmt.Errorf("execute if expression command part failed due to %s", err.Error())
	}
	goThen, err := data.CoerceToBoolean(exprValue)
	if err != nil {
		return nil, fmt.Errorf("if expression must be an boolean, convert to boolean failed due to %s", err.Error())
	}

	if goThen {
		//Then part
		return iff.Then.EvalWithScope(inputScope, resolver)
	} else {
		//else part
		if iff.Else != nil {
			return iff.Else.EvalWithScope(inputScope, resolver)
		}
		return nil, nil
	}
}

func (iff *iff) Eval() (interface{}, error) {
	exprValue, err := iff.Expr.Eval()
	if err != nil {
		return nil, fmt.Errorf("execute if expression command part failed due to %s", err.Error())
	}
	goThen, err := data.CoerceToBoolean(exprValue)
	if err != nil {
		return nil, fmt.Errorf("if expression must be an boolean, convert to boolean failed due to %s", err.Error())
	}

	if goThen {
		//Then part
		return iff.Then.Eval()
	} else {
		//else part
		if iff.Else != nil {
			return iff.Else.Eval()
		}
		return nil, nil
	}
}

func (iff *iff) EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	exprValue, err := iff.Expr.EvalWithData(value, inputScope, resolver)
	if err != nil {
		return nil, fmt.Errorf("execute if expression command part failed due to %s", err.Error())
	}
	goThen, err := data.CoerceToBoolean(exprValue)
	if err != nil {
		return nil, fmt.Errorf("if expression must be an boolean, convert to boolean failed due to %s", err.Error())
	}

	if goThen {
		//Then part
		return iff.Then.EvalWithData(value, inputScope, resolver)
	} else {
		//else part
		if iff.Else != nil {
			return iff.Else.EvalWithData(value, inputScope, resolver)
		}
		return nil, nil
	}
}

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
	return iffl.V, nil

}
