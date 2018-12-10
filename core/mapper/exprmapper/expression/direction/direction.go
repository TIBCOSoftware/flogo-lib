package direction

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/iff"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"reflect"
	"strconv"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/data"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/expr"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/gocc/token"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/ref"
)

var log = logger.GetLogger("expression-direction")

type Attribute interface{}

func NewDoubleQuoteStringLit(lit interface{}) (string, error) {
	str := strings.TrimSpace(string(lit.(*token.Token).Lit))

	if str != "" && len(str) > 0 {
		str = RemoveQuote(str)
	}
	//Eascap string
	if strings.Contains(str, "\\\"") {
		str = strings.Replace(str, "\\\"", "\"", -1)
	}

	return str, nil
}

func NewSingleQuoteStringLit(lit interface{}) (string, error) {
	str := strings.TrimSpace(string(lit.(*token.Token).Lit))

	if str != "" && len(str) > 0 {
		str = RemoveQuote(str)
	}

	//Eascap string
	if strings.Contains(str, "\\'") {
		str = strings.Replace(str, "\\'", "'", -1)
	}

	return str, nil
}

func NewIntLit(lit interface{}) (int, error) {
	str := strings.TrimSpace(string(lit.(*token.Token).Lit))
	s, err := data.CoerceToInteger(str)
	return s, err
}

func NewFloatLit(lit interface{}) (float64, error) {
	str := strings.TrimSpace(string(lit.(*token.Token).Lit))
	s, err := data.CoerceToNumber(str)
	return s, err
}

func NewBool(lit interface{}) (bool, error) {
	s := strings.TrimSpace(string(lit.(*token.Token).Lit))
	b, err := strconv.ParseBool(s)
	return b, err
}

type NIL struct {
}

func NewNilLit(lit interface{}) (*NIL, error) {
	return &NIL{}, nil
}

func NewMappingRef(lit interface{}) (interface{}, error) {
	s := strings.TrimSpace(string(lit.(*token.Token).Lit))
	if strings.HasPrefix(s, "$.") || strings.HasPrefix(s, "$$") {
		m := ref.NewArrayRef(s)
		return m, nil
	} else {
		m := ref.NewMappingRef(s)
		return m, nil
	}
}

func NewFunction(name Attribute, parameters Attribute) (interface{}, error) {
	f_func := &function.FunctionExp{}
	to := name.(*token.Token)
	f_func.Name = string(to.Lit)

	switch parameters.(type) {
	case *function.Parameter:
		f_func.Params = append(f_func.Params, parameters.(*function.Parameter))
	case []*function.Parameter:
		for _, p := range parameters.([]*function.Parameter) {
			f_func.Params = append(f_func.Params, p)
		}
	}

	return f_func, nil
}

func NewArgument(a Attribute) (interface{}, error) {
	parameters := []*function.Parameter{}
	switch a.(type) {
	case *NIL:
		param := &function.Parameter{Value: expr.NewLiteralExpr(nil)}
		parameters = append(parameters, param)
	case expr.Expr:
		param := &function.Parameter{Value: a.(expr.Expr)}
		parameters = append(parameters, param)
	case []*function.Parameter:
		for _, p := range a.([]*function.Parameter) {
			parameters = append(parameters, p)
		}
	case []interface{}:
		//TODO
		log.Debug("New Arguments type is []interface{}")
	case interface{}:
		//TODO
		log.Debugf("New Arguments type is interface{}, [%+v]", reflect.TypeOf(a))
	}
	return parameters, nil
}

func NewArguments(as ...Attribute) (interface{}, error) {
	parameters := []*function.Parameter{}
	for _, a := range as {
		switch a.(type) {
		case *NIL:
			param := &function.Parameter{Value: expr.NewLiteralExpr(nil)}
			parameters = append(parameters, param)
		case expr.Expr:
			param := &function.Parameter{Value: a.(expr.Expr)}
			parameters = append(parameters, param)
		case []*function.Parameter:
			for _, p := range a.([]*function.Parameter) {
				parameters = append(parameters, p)
			}
		default:
			param := &function.Parameter{Value: expr.NewLiteralExpr(a)}
			parameters = append(parameters, param)
		}
	}
	return parameters, nil
}

func NewExpressionField(a Attribute) (interface{}, error) {
	expression := getExpression(a)
	return expression, nil
}

func NewLiteralExpr(a Attribute) (interface{}, error) {
	var literalExpr expr.Expr
	switch t := a.(type) {
	case expr.Expr:
		literalExpr = t
	case *NIL:
		literalExpr = expr.NewLiteralExpr(nil)
	default:
		literalExpr = expr.NewLiteralExpr(t)
	}
	return literalExpr, nil
}

func NewExpression(left Attribute, op Attribute, right Attribute) (interface{}, error) {
	expression := expr.NewExpression()
	operator := strings.TrimSpace(string(op.(*token.Token).Lit))
	expression.Operator = operator

	expression.Left = getExpression(left)
	expression.Right = getExpression(right)
	log.Debugf("New expression left [%+v] right [%s+v and operator [%s]", expression.Left, expression.Right, operator)
	return expression, nil
}

func getExpression(ex Attribute) *expr.Expression {
	expression := expr.NewExpression()
	switch ex.(type) {
	case expr.Expr:
		expression.Value = ex.(expr.Expr)
	default:
		expression.Value = expr.NewLiteralExpr(ex)
	}
	return expression
}

func NewTernaryExpression(first Attribute, second Attribute, third Attribute) (Attribute, error) {
	log.Debugf("first [%+v] and type [%s]", first, reflect.TypeOf(first))
	log.Debugf("second [%+v] and type [%s]", second, reflect.TypeOf(second))
	log.Debugf("third [%+v] and type [%s]", third, reflect.TypeOf(third))
	ternaryExp := &expr.TernaryExpressio{First: first, Second: second, Third: third}
	return ternaryExp, nil
}

func NewIFF(ifExpr Attribute, statement Attribute) (interface{}, error) {
	var ifE, statementE expr.Expr
	switch t := ifExpr.(type) {
	case expr.Expr:
		ifE = t
	default:
		ifE = expr.NewIfValueProvider(t)
	}

	switch t := statement.(type) {
	case expr.Expr:
		statementE = t
	default:
		statementE = expr.NewIfValueProvider(t)
	}

	return iff.NewIff(ifE, statementE), nil
}

func NewStatement(statement Attribute) (interface{}, error) {
	var statementE expr.Expr

	switch t := statement.(type) {
	case expr.Expr:
		statementE = t
	default:
		statementE = expr.NewIfValueProvider(t)
	}

	return statementE, nil
}

func NewIFFElse(ifExpr Attribute, statement, elseS Attribute) (interface{}, error) {
	var ifE, statementE, elseE expr.Expr
	switch t := ifExpr.(type) {
	case expr.Expr:
		ifE = t
	default:
		ifE = expr.NewIfValueProvider(t)
	}

	switch t := statement.(type) {
	case expr.Expr:
		statementE = t
	default:
		statementE = expr.NewIfValueProvider(t)
	}

	switch t := elseS.(type) {
	case expr.Expr:
		elseE = t
	default:
		elseE = expr.NewIfValueProvider(t)

	}

	return iff.NewIffElse(ifE, statementE, elseE), nil
}

func RemoveQuote(quoteStr string) string {
	if HasQuote(quoteStr) {
		if strings.HasPrefix(quoteStr, `"`) || strings.HasPrefix(quoteStr, `'`) {
			quoteStr = quoteStr[1 : len(quoteStr)-1]
		}
	}
	return quoteStr
}

func HasQuote(quoteStr string) bool {
	if strings.HasPrefix(quoteStr, `"`) && strings.HasSuffix(quoteStr, `"`) {
		return true
	}

	if strings.HasPrefix(quoteStr, `'`) && strings.HasSuffix(quoteStr, `'`) {
		return true
	}

	return false
}
