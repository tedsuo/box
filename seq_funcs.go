package box

import (
	"fmt"
	"reflect"
)

func Concat(args ...interface{}) (out Sequence) {
	out = NewSeq()
	go func() {
		defer close(out)
		for _, arg := range args {
			for ß := range NewSeq(arg) {
				out <- ß
			}
		}
	}()
	return
}

func Count(input interface{}) int {
	i := 0
	Each(input, func() {
		i++
	})
	return i
}

func Each(sequenceInput interface{}, callbacks ...interface{}) {
	cbStack := newCallbackStack(callbacks)
	for ß := range NewSeq(sequenceInput) {
		for _, cb := range cbStack {
			cb.Call(ß)
		}
	}
}

func Stream(sequenceInput interface{}, callbacks ...interface{}) (seq Sequence) {
	seq = NewSeq()
	cbStack := newCallbackStack(callbacks)
	go func() {
		defer close(seq)
		for boxedVal := range NewSeq(sequenceInput) {
			for _, cb := range cbStack {
				for ß := range NewSeq(cb.Call(boxedVal)) {
					seq <- ß
				}
			}
		}
	}()
	return
}

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

var emptyArgs = []reflect.Value{}

func (cb *callback) Call(ß Box) (results []interface{}) {
	results = []interface{}{}
	if cb.argLength == 0 {
		cb.funcVal.Call(emptyArgs)
		return
	}

	var args = []reflect.Value{}
	for i := len(ß)-cb.argLength; i < len(ß); i++ {
		boxedVal := ß[i]
		args = append(args, reflect.ValueOf(boxedVal))
	}

	resultVals := cb.funcVal.Call(args)
	for _, resultVal := range resultVals {
		results = append(results, resultVal.Interface())
	}
	return
}
