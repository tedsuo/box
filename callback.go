package box

import (
	"fmt"
	"reflect"
)

var emptyArgs = []reflect.Value{}

type callbackStack []*callback

func newCallbackStack(callbacks []interface{}) (stack callbackStack) {
	stack = callbackStack{}
	for _, callback := range callbacks {
		cb := newCallback(callback)
		stack = append(stack, cb)
	}
	return
}

type callback struct {
	funcVal   reflect.Value
	argLength int
}

func newCallback(callbackInput interface{}) *callback {
	cb := reflect.ValueOf(callbackInput)
	call := new(callback)
	call.funcVal = cb
	call.argLength = cb.Type().NumIn()
	return call
}

func (cb *callback) String() string {
	return fmt.Sprintf("%+v", *cb)
}

func (cb *callback) Call(ß Box) (results []interface{}) {
	results = []interface{}{}
	if cb.argLength == 0 {
		cb.funcVal.Call(emptyArgs)
		return
	}

	var args = []reflect.Value{}
	for i := len(ß) - cb.argLength; i < len(ß); i++ {
		boxedVal := ß[i]
		args = append(args, reflect.ValueOf(boxedVal))
	}

	resultVals := cb.funcVal.Call(args)
	for _, resultVal := range resultVals {
		results = append(results, resultVal.Interface())
	}
	return
}
