package function

import (
	"errors"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/expr"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"reflect"
	"runtime/debug"
)

var logrus = logger.GetLogger("function")

type Func interface {
	Eval(inputScope, outputScope data.Scope) ([]interface{}, error)
	String() string
}

type FunctionExp struct {
	Name   string       `json:"name"`
	Params []*Parameter `json:"params"`
}

type Parameter struct {
	Value expr.Expr `json:"value"`
}

func (f *FunctionExp) Eval() (interface{}, error) {
	value, err := f.callFunction(nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return convertType(value), err
}

func (f *FunctionExp) EvalWithScope(inputScope data.Scope, resolver data.Resolver) (interface{}, error) {

	value, err := f.callFunction(nil, inputScope, resolver)
	if err != nil {
		logrus.Errorf("Execution failed for function [%s] error - %+v", f.Name, err.Error())
		return nil, err
	}
	return convertType(value), err
}

func (f *FunctionExp) EvalWithData(data interface{}, inputScope data.Scope, resolver data.Resolver) (interface{}, error) {
	value, err := f.callFunction(data, inputScope, resolver)
	if err != nil {
		return nil, err
	}
	return convertType(value), err
}

func HandleToSingleOutput(values interface{}) interface{} {
	if values != nil {
		switch t := values.(type) {
		case []interface{}:
			return t[0]
		default:
			return t
		}
	}
	return nil
}

func convertType(value reflect.Value) interface{} {
	return value.Interface()
}

func (f *FunctionExp) getRealFunction() (Function, error) {
	return GetFunction(f.Name)
}

func (f *FunctionExp) getMethod() (reflect.Value, error) {
	var ptr reflect.Value
	s, err := f.getRealFunction()
	if err != nil {
		return reflect.Value{}, err
	}

	value := reflect.ValueOf(s)
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(s))
		temp := ptr.Elem()
		temp.Set(value)
	}

	method := value.MethodByName("Eval")
	if !method.IsValid() {
		method = ptr.MethodByName("Eval")
		if !method.IsValid() {
			logrus.Debug("invalid also, ", f.Name)
			return reflect.Value{}, errors.New("Method invalid..")

		}
	}

	return method, nil
}

func (f *FunctionExp) callFunction(fdata interface{}, inputScope data.Scope, resolver data.Resolver) (results reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%+v", r)
			logrus.Debugf("StackTrace: %s", debug.Stack())
		}
	}()

	method, err := f.getMethod()
	if err != nil {
		return reflect.Value{}, err
	}

	inputs := []reflect.Value{}
	for _, p := range f.Params {
		result, err := p.Value.EvalWithData(fdata, inputScope, resolver)
		if err != nil {
			return reflect.Value{}, err
		}

		logrus.Debugf("function [%s] [%d]'s argument value [%+v]", f.Name, result)
		inputs = append(inputs, reflect.ValueOf(result))
	}

	logrus.Debugf("Input Parameters: %+v", inputs)
	args, err := ensureArguments(method, inputs)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("Function '%s' argument validation failed due to error %s", f.Name, err.Error())
	}
	values := method.Call(args)
	return f.extractErrorFromValues(values)
}

func ensureArguments(method reflect.Value, in []reflect.Value) ([]reflect.Value, error) {

	var retInputs []reflect.Value
	methodType := method.Type()
	n := method.Type().NumIn()

	if methodType.IsVariadic() && n == 1 {
		x := in[0]
		elem := methodType.In(0).Elem()
		if xt := x.Type(); !xt.AssignableTo(elem) {
			v, err := convertArgs(elem, x)
			if err != nil {
				return nil, fmt.Errorf("argument type mismatch. Can not convert type %s to type %s. ", xt.String(), elem.String())
			}
			retInputs = append(retInputs, reflect.ValueOf(v))
		} else {
			retInputs = append(retInputs, x)
		}
	} else {
		for i := 0; i < n; i++ {
			if xt, targ := in[i].Type(), methodType.In(i); !xt.AssignableTo(targ) {
				v, err := convertArgs(targ, in[i])
				if err != nil {
					return nil, fmt.Errorf("argument type mismatch. Can not convert type %s to type %s. ", xt.String(), targ.String())
				}
				retInputs = append(retInputs, reflect.ValueOf(v))
			} else {
				retInputs = append(retInputs, in[i])
			}
		}
	}

	if methodType.IsVariadic() {
		m := len(in) - n
		elem := methodType.In(n - 1).Elem()
		for j := 0; j < m; j++ {
			x := in[n+j]
			emtpy := reflect.Value{}
			if x == emtpy {
				retInputs = append(retInputs, reflect.Zero(elem))
			} else {
				if xt := x.Type(); !xt.AssignableTo(elem) {
					v, err := convertArgs(elem, x)
					if err != nil {
						return nil, fmt.Errorf("argument type mismatch. Can not convert type %s to type %s. ", xt.String(), elem.String())
					}
					retInputs = append(retInputs, reflect.ValueOf(v))
				} else {
					retInputs = append(retInputs, x)
				}
			}

		}
	}

	return retInputs, nil
}

func convertArgs(argType reflect.Type, in reflect.Value) (interface{}, error) {
	var v interface{}
	var err error
	switch argType.Kind() {
	case reflect.Bool:
		v, err = data.CoerceToBoolean(in.Interface())
	case reflect.Interface:
		v, err = data.CoerceToAny(in.Interface())
	case reflect.Int:
		v, err = data.CoerceToInteger(in.Interface())
	case reflect.Int64:
		v, err = data.CoerceToLong(in.Interface())
	case reflect.String:
		v, err = data.CoerceToString(in.Interface())
	case reflect.Float64:
		v, err = data.CoerceToDouble(in.Interface())
	default:
		v = in.Interface()
	}
	return v, err

}

func (f *FunctionExp) extractErrorFromValues(values []reflect.Value) (reflect.Value, error) {
	tempValues := []reflect.Value{}

	var err error
	for _, value := range values {
		if value.Type().Name() == "error" {
			if value.Interface() != nil {
				err = value.Interface().(error)
			}
		} else {
			tempValues = append(tempValues, value)
		}
	}

	if len(tempValues) > 1 {
		return tempValues[0], fmt.Errorf("Not support function multiple returns")
	}
	return tempValues[0], err
}
