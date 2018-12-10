package expr

import (
	"encoding/json"
	"errors"
	"reflect"

	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/ref"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("expr")

const (
	OR             = "||"
	AND            = "&&"
	EQ             = "=="
	NOT_EQ         = "!="
	GT             = ">"
	LT             = "<"
	GTE            = ">="
	LTE            = "<="
	ADDITION       = "+"
	SUBTRACTION    = "-"
	MULTIPLICATION = "*"
	DIVIDE         = "/"
	MODE           = "%"
)

type Expr interface {
	EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error)
	Eval() (interface{}, error)
	EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error)
}

type Expression struct {
	Left     Expr   `json:"left"`
	Operator string `json:"operator"`
	Right    Expr   `json:"right"`
	Value    Expr   `json:"value"`
}

func (e *Expression) IsNil() bool {
	if e.Left == nil && e.Right == nil {
		return true
	}
	return false
}

type TernaryExpressio struct {
	First  interface{}
	Second interface{}
	Third  interface{}
}

func (t *TernaryExpressio) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	return t.EvalWithData(nil, inputScope, resolver)
}

func (t *TernaryExpressio) Eval() (interface{}, error) {
	return t.EvalWithScope(nil, data.GetBasicResolver())
}

func (t *TernaryExpressio) EvalWithData(value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	v, err := t.HandleParameter(t.First, value, inputScope, resolver)
	if err != nil {
		return nil, err
	}
	if v.(bool) {
		v2, err2 := t.HandleParameter(t.Second, value, inputScope, resolver)
		if err2 != nil {
			return nil, err2
		}
		return v2, nil
	} else {
		v3, err3 := t.HandleParameter(t.Third, value, inputScope, resolver)
		if err3 != nil {
			return nil, err3
		}
		return v3, nil
	}
}

func (t *TernaryExpressio) HandleParameter(param interface{}, value interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	var firstValue interface{}
	switch t := param.(type) {
	case Expr:
		vss, err := t.EvalWithData(value, inputScope, resolver)
		if err != nil {
			return nil, err
		}
		firstValue = vss
		return firstValue, nil
	case *ref.ArrayRef:
		return handleArrayRef(value, t.GetRef(), inputScope, resolver)
	case *ref.MappingRef:
		return t.Eval(inputScope, resolver)
	default:
		firstValue = t
		return firstValue, nil
	}
}

func handleArrayRef(edata interface{}, mapref string, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	if edata == nil {
		v, err := ref.NewMappingRef(mapref).Eval(inputScope, resolver)
		if err != nil {
			log.Errorf("Mapping ref eva error [%s]", err.Error())
			return nil, fmt.Errorf("Mapping ref eva error [%s]", err.Error())
		}
		return v, nil
	} else {
		arrayRef := ref.NewArrayRef(mapref)
		v, err := arrayRef.EvalFromData(edata)
		if err != nil {
			log.Errorf("Mapping ref eva error [%s]", err.Error())
			return nil, fmt.Errorf("Mapping ref eva error [%s]", err.Error())
		}
		return v, nil
	}
}

func (e *Expression) String() string {
	v, err := json.Marshal(e)
	if err != nil {
		log.Errorf("Expression to string error [%s]", err.Error())
		return ""
	}
	return string(v)
}

func NewExpression() *Expression {
	return &Expression{}
}

func (f *Expression) Eval() (interface{}, error) {
	return f.evaluate(nil, nil, nil)
}

func (f *Expression) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	return f.evaluate(nil, inputScope, resolver)
}

func (f *Expression) EvalWithData(data interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	return f.evaluate(data, inputScope, resolver)
}

func (f *Expression) evaluate(data interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	//Left
	if f.IsNil() {
		log.Debugf("Expression right and left are nil, return value directly")
		return f.Value.EvalWithData(data, inputScope, resolver)
	}

	var leftValue interface{}
	var rightValue interface{}

	if f.Left != nil {
		leftResultChan := make(chan interface{}, 1)
		go do(f.Left, data, inputScope, resolver, leftResultChan)
		leftValue = <-leftResultChan
	}

	if f.Right != nil {
		rightResultChan := make(chan interface{}, 1)
		go do(f.Right, data, inputScope, resolver, rightResultChan)
		rightValue = <-rightResultChan
	}

	//Make sure no error returned
	switch leftValue.(type) {
	case error:
		return nil, leftValue.(error)
	}

	switch rightValue.(type) {
	case error:
		return nil, rightValue.(error)
	}
	//Operator
	operator := f.Operator

	return f.run(leftValue, operator, rightValue)
}

func do(f Expr, edata interface{}, inputScope data.Scope, resolver data.Resolver, resultChan chan interface{}) {
	if f == nil {
		resultChan <- nil
	}
	leftValue, err := f.EvalWithData(edata, inputScope, resolver)
	if err != nil {
		resultChan <- errors.New("Eval left expression error: " + err.Error())
	}
	resultChan <- leftValue
}

func (f *Expression) run(left interface{}, op string, right interface{}) (interface{}, error) {
	switch op {
	case EQ:
		return equals(left, right)
	case OR:
		return or(left, right)
	case AND:
		return add(left, right)
	case NOT_EQ:
		return notEquals(left, right)
	case GT:
		return gt(left, right, false)
	case LT:
		return lt(left, right, false)
	case GTE:
		return gt(left, right, true)
	case LTE:
		return lt(left, right, true)
	case ADDITION:
		return additon(left, right)
	case SUBTRACTION:
		return sub(left, right)
	case MULTIPLICATION:
		return multiplication(left, right)
	case DIVIDE:
		return div(left, right)
	case MODE:
		return mod(left, right)
	default:
		return nil, errors.New("Unknow operator " + op)
	}

	return nil, nil

}

func equals(left interface{}, right interface{}) (bool, error) {
	log.Debugf("Equals condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return true, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	leftValue, rightValue, err := ConvertToSameType(left, right)
	if err != nil {
		return false, err
	}

	log.Debugf("Equals condition -> right expression value [%s]", rightValue)

	return leftValue == rightValue, nil
}

func ConvertToSameType(left interface{}, right interface{}) (interface{}, interface{}, error) {
	if left == nil || right == nil {
		return left, right, nil
	}
	var leftValue interface{}
	var rightValue interface{}
	var err error
	switch t := left.(type) {
	case int:
		if isDoubleType(right) {
			leftValue, err = data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue = right
		} else {
			rightValue, err = data.CoerceToInteger(right)
			if err != nil {
				err = fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			leftValue = t
		}

	case int64:
		if isDoubleType(right) {
			leftValue, err = data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue = right
		} else {
			rightValue, err = data.CoerceToInteger(right)
			if err != nil {
				err = fmt.Errorf("Convert right expression to type int64 failed, due to %s", err.Error())
			}
			leftValue = t
		}
	case float64:
		rightValue, err = data.CoerceToNumber(right)
		if err != nil {
			err = fmt.Errorf("Convert right expression to type float64 failed, due to %s", err.Error())
		}
		leftValue = t
	case string:
		rightValue, err = data.CoerceToString(right)
		if err != nil {
			err = fmt.Errorf("Convert right expression to type string failed, due to %s", err.Error())
		}
		leftValue = t

	case bool:
		rightValue, err = data.CoerceToBoolean(right)
		if err != nil {
			err = fmt.Errorf("Convert right expression to type boolean failed, due to %s", err.Error())
		}
		leftValue = t

	case json.Number:
		rightValue, err = data.CoerceToLong(right)
		if err != nil {
			err = fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err = data.CoerceToLong(left)
		if err != nil {
			err = fmt.Errorf("Convert left expression to type long failed, due to %s", err.Error())
		}
	default:
		err = fmt.Errorf("Unsupport type to compare now")
	}
	return leftValue, rightValue, err

}

func notEquals(left interface{}, right interface{}) (bool, error) {

	log.Debugf("Not equals condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return true, nil
	} else if left != nil && right == nil {
		return true, nil
	}

	leftValue, rightValue, err := ConvertToSameType(left, right)
	if err != nil {
		return false, err
	}

	log.Debugf("Not equals condition -> right expression value [%s]", rightValue)

	return leftValue != rightValue, nil

}

func gt(left interface{}, right interface{}, includeEquals bool) (bool, error) {

	log.Debugf("Greater than condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	log.Debugf("Greater than condition -> right value [%+v] and Right value: [%+v]", left, right)
	rightType := getType(right)
	switch le := left.(type) {
	case int:
		//For int float compare, convert int to float to compare
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			if includeEquals {
				return leftValue >= rightValue, nil

			} else {
				return leftValue > rightValue, nil
			}

		} else {
			//We should conver to int first
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			if includeEquals {
				return le >= rightValue, nil

			} else {
				return le > rightValue, nil
			}
		}

	case int64:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			if includeEquals {
				return leftValue >= rightValue, nil

			} else {
				return leftValue > rightValue, nil
			}

		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			if includeEquals {
				return int(le) >= rightValue, nil

			} else {
				return int(le) > rightValue, nil
			}
		}
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		if includeEquals {
			return le >= rightValue, nil

		} else {
			return le > rightValue, nil
		}
	case string, json.Number:
		//In case of string, convert to number and compare
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int64 failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert left expression to type int64 failed, due to %s", err.Error())
		}

		if includeEquals {
			return leftValue >= rightValue, nil

		} else {
			return leftValue > rightValue, nil
		}
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to greater than, left [%s] and right [%s] ", getType(left).String(), rightType.String()))
	}

	return false, nil
}

func lt(left interface{}, right interface{}, includeEquals bool) (bool, error) {

	log.Debugf("Less than condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	switch le := left.(type) {
	case int:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			if includeEquals {
				return leftValue <= rightValue, nil

			} else {
				return leftValue < rightValue, nil
			}
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			if includeEquals {
				return le <= rightValue, nil

			} else {
				return le < rightValue, nil
			}
		}
	case int64:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			if includeEquals {
				return leftValue <= rightValue, nil

			} else {
				return leftValue < rightValue, nil
			}
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			if includeEquals {
				return int(le) <= rightValue, nil

			} else {
				return int(le) < rightValue, nil
			}
		}
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		if includeEquals {
			return le <= rightValue, nil

		} else {
			return le < rightValue, nil
		}
	case string, json.Number:
		//In case of string, convert to number and compare
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int64 failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert left expression to type int64 failed, due to %s", err.Error())
		}

		if includeEquals {
			return leftValue <= rightValue, nil

		} else {
			return leftValue < rightValue, nil
		}
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to <, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func add(left interface{}, right interface{}) (bool, error) {

	log.Debugf("Add condition -> left expression value %+v, right expression value %+v", left, right)

	switch le := left.(type) {
	case bool:
		rightValue, err := data.CoerceToBoolean(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le && rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to &&, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func or(left interface{}, right interface{}) (bool, error) {

	log.Debugf("Or condition -> left expression value %+v, right expression value %+v", left, right)
	switch le := left.(type) {
	case bool:
		rightValue, err := data.CoerceToBoolean(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le || rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to ||, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func additon(left interface{}, right interface{}) (interface{}, error) {

	log.Debugf("Addition condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	switch le := left.(type) {
	case int:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue + rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}

			return le + rightValue, nil
		}

	case int64:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue + rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			return int(le) + rightValue, nil
		}
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le + rightValue, nil
	case json.Number:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		return leftValue * rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to additon, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func sub(left interface{}, right interface{}) (interface{}, error) {

	log.Debugf("Sub condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	switch le := left.(type) {
	case int:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue - rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			return le - rightValue, nil
		}
	case int64:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue - rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			return int(le) - rightValue, nil
		}
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}

		return le - rightValue, nil
	case json.Number:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		return leftValue - rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to sub, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func multiplication(left interface{}, right interface{}) (interface{}, error) {

	log.Debugf("Multiplication condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil && right == nil {
		return false, nil
	} else if left == nil && right != nil {
		return false, nil
	} else if left != nil && right == nil {
		return false, nil
	}

	switch le := left.(type) {
	case int:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue * rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}
			return le * rightValue, nil
		}
	case int64:
		if isDoubleType(right) {
			leftValue, err := data.CoerceToDouble(left)
			if err != nil {
				err = fmt.Errorf("Convert left expression to type float64 failed, due to %s", err.Error())
			}
			rightValue := right.(float64)
			return leftValue * rightValue, nil
		} else {
			rightValue, err := data.CoerceToInteger(right)
			if err != nil {
				return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
			}

			return int(le) * rightValue, nil
		}
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}

		return le * rightValue, nil
	case json.Number:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		return leftValue * rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to multiplication, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func div(left interface{}, right interface{}) (interface{}, error) {

	log.Debugf("Div condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil || right == nil {
		return nil, fmt.Errorf("Cannot run dividing operation on empty value")
	}

	switch le := left.(type) {
	case int:
		rightValue, err := data.CoerceToInteger(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le / rightValue, nil
	case int64:
		rightValue, err := data.CoerceToInteger(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return int(le) / rightValue, nil
	case float64:
		rightValue, err := data.CoerceToNumber(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le / rightValue, nil
	case json.Number:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		return leftValue / rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to div, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func mod(left interface{}, right interface{}) (interface{}, error) {

	log.Debugf("% condition -> left expression value %+v, right expression value %+v", left, right)
	if left == nil || right == nil {
		return nil, fmt.Errorf("Cannot run mod operation on empty value")
	}

	switch le := left.(type) {
	case int:
		rightValue, err := data.CoerceToInteger(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return le % rightValue, nil
	case int64:
		rightValue, err := data.CoerceToInteger(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		return int(le) % rightValue, nil
	case float64:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type int failed, due to %s", err.Error())
		}
		lev := int64(le)
		return lev % rightValue, nil
	case json.Number:
		rightValue, err := data.CoerceToLong(right)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		leftValue, err := data.CoerceToLong(left)
		if err != nil {
			return false, fmt.Errorf("Convert right expression to type long failed, due to %s", err.Error())
		}

		return leftValue % rightValue, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknow type use to div, left [%s] and right [%s] ", getType(left).String(), getType(right).String()))
	}

	return false, nil
}

func getType(in interface{}) reflect.Type {
	return reflect.TypeOf(in)
}

func isDoubleType(in interface{}) bool {
	switch in.(type) {
	case float64:
		return true
	}
	return false
}
