package direction

import (
	"errors"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/iff"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"reflect"
	"strconv"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/data"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/expr"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/gocc/token"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/funcexprtype"
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
			if !p.IsEmtpy() {
				f_func.Params = append(f_func.Params, p)
			}
		}
	}

	log.Debug(f_func.Tostring())
	return f_func, nil
}

func NewArgument(a Attribute) (interface{}, error) {
	param := &function.Parameter{}
	parameters := []*function.Parameter{}
	switch t := a.(type) {
	case *token.Token:
		param.Type = funcexprtype.STRING
		param.Value = string(a.(*token.Token).Lit)
	case int:
		param.Type = funcexprtype.INTEGER
		param.Value = a
	case float64:
		param.Type = funcexprtype.FLOAT
		param.Value = a
	case string:
		param.Type = funcexprtype.STRING
		param.Value = a.(string)
	case bool:
		param.Type = funcexprtype.BOOLEAN
		param.Value = a.(bool)
	case *NIL:
		param.Type = funcexprtype.NIL
		param.Value = nil
	case *function.FunctionExp:
		param.Type = funcexprtype.FUNCTION
		param.Function = a.(*function.FunctionExp)
	case []*function.Parameter:
		for _, p := range a.([]*function.Parameter) {
			if !p.IsEmtpy() {
				parameters = append(parameters, p)
			}
		}
	case *expr.Expression:
		exprFieldToArgument(t, param)
	case *ref.MappingRef:
		param.Type = funcexprtype.REF
		param.Value = a
	case *ref.ArrayRef:
		param.Type = funcexprtype.ARRAYREF
		param.Value = a
	case []interface{}:
		//TODO
		log.Debug("New Arguments type is []interface{}")
	case interface{}:
		//TODO
		log.Debugf("New Arguments type is interface{}, [%+v]", reflect.TypeOf(a))
	}
	parameters = append(parameters, param)
	return parameters, nil
}

func exprFieldToArgument(ex *expr.Expression, param *function.Parameter) {
	if ex != nil {
		switch ex.Type {
		case funcexprtype.INTEGER, funcexprtype.ARRAYREF, funcexprtype.BOOLEAN, funcexprtype.FLOAT, funcexprtype.REF, funcexprtype.STRING:
			param.Type = ex.Type
			param.Value = ex.Value
		case funcexprtype.FUNCTION:
			param.Type = ex.Type
			param.Function = ex.Value.(*function.FunctionExp)
		}
	}
}

func NewArguments(as ...Attribute) (interface{}, error) {
	parameters := []*function.Parameter{}
	for _, a := range as {
		param := &function.Parameter{}
		switch a.(type) {
		case *token.Token:
			param.Type = funcexprtype.STRING
			param.Value = string(a.(*token.Token).Lit)
		case int:
			param.Type = funcexprtype.INTEGER
			param.Value = a
		case float64:
			param.Type = funcexprtype.FLOAT
			param.Value = a
		case string:
			param.Type = funcexprtype.STRING
			param.Value = a.(string)
		case *function.FunctionExp:
			param.Type = funcexprtype.FUNCTION
			param.Function = a.(*function.FunctionExp)
		case *ref.MappingRef:
			param.Type = funcexprtype.REF
			param.Value = a
		case *ref.ArrayRef:
			param.Type = funcexprtype.ARRAYREF
			param.Value = a
		case []*function.Parameter:
			for _, p := range a.([]*function.Parameter) {
				if !p.IsEmtpy() {
					parameters = append(parameters, p)
				}
			}
		case []interface{}:
			log.Debugf("New Arguments type is []interface{}")
		case interface{}:
			log.Debugf("New Arguments type is interface{} %+v", a)
		}
		parameters = append(parameters, param)
	}
	return parameters, nil
}

func NewExpressionField(a Attribute) (interface{}, error) {
	expression := getExpression(a)
	return expression, nil
}

func NewExpression(left Attribute, op Attribute, right Attribute) (interface{}, error) {
	expression := expr.NewExpression()
	operator, found := expr.ToOperator(strings.TrimSpace(string(op.(*token.Token).Lit)))
	if found {
		expression.Operator = operator
	} else {
		return nil, errors.New("Unsupport operator " + string(op.(*token.Token).Lit))
	}

	expression.Left = getExpression(left)
	expression.Right = getExpression(right)
	expression.Type = funcexprtype.EXPRESSION
	log.Debugf("New expression left [%+v] right [%s+v and operator [%s]", expression.Left, expression.Right, expression.Operator)
	return expression, nil
}

func getExpression(ex Attribute) *expr.Expression {
	expression := expr.NewExpression()
	switch ex.(type) {
	case int:
		expression.Type = funcexprtype.INTEGER
		expression.Value = ex.(int)
	case float64:
		expression.Type = funcexprtype.FLOAT
		expression.Value = ex.(float64)
	case string:
		expression.Type = funcexprtype.STRING
		expression.Value = ex.(string)
	case bool:
		expression.Type = funcexprtype.BOOLEAN
		expression.Value = ex.(bool)
	case *NIL:
		expression.Type = funcexprtype.NIL
		expression.Value = nil
	case ref.MappingRef:
		expression.Type = funcexprtype.REF
		ref := ex.(ref.MappingRef)
		expression.Value = ref
	case *ref.MappingRef:
		expression.Type = funcexprtype.REF
		expression.Value = ex.(*ref.MappingRef).GetRef()
	case *ref.ArrayRef:
		expression.Type = funcexprtype.ARRAYREF
		expression.Value = ex.(*ref.ArrayRef).GetRef()
	case *function.FunctionExp:
		expression.Type = funcexprtype.FUNCTION
		expression.Value = ex.(*function.FunctionExp)
	case *expr.Expression:
		expression = ex.(*expr.Expression)
	default:
		log.Errorf("Unknow expression type [%s]", ex)
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
		ifE = iff.NewIffExpr(t)
	}

	switch t := statement.(type) {
	case expr.Expr:
		statementE = t
	default:
		statementE = iff.NewIffExpr(t)
	}

	return iff.NewIff(ifE, statementE), nil
}

func NewIFFElse(ifExpr Attribute, statement, elseS Attribute) (interface{}, error) {
	var ifE, statementE, elseE expr.Expr
	switch t := ifExpr.(type) {
	case expr.Expr:
		ifE = t
	default:
		ifE = iff.NewIffExpr(t)
	}

	switch t := statement.(type) {
	case expr.Expr:
		statementE = t
	default:
		statementE = iff.NewIffExpr(t)
	}

	switch t := elseS.(type) {
	case expr.Expr:
		elseE = t
	default:
		elseE = iff.NewIffExpr(t)

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
