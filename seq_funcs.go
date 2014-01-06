package box

import (
	"fmt"
	"reflect"
)

func Concat(args ...interface{}) (out Sequence) {
	out = NewSeq()
	go func() {
		defer close(out)
		for seq := range NewSeq(args) {
			for ß := range NewSeq(seq) {
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
	for boxedVal := range NewSeq(sequenceInput) {
		for _, cb := range cbStack {
			cb.Call(boxedVal)
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
	args      []reflect.Value
}

func newCallback(callbackInput interface{}) *callback {
	cb := reflect.ValueOf(callbackInput)
	call := new(callback)
	call.funcVal = cb
	call.argLength = cb.Type().NumIn()
	call.args = []reflect.Value{}
	return call
}

func (cb *callback) String() string {
	return fmt.Sprintf("%+v", *cb)
}

func (cb *callback) Call(args ...interface{}) (results []interface{}) {
	results = []interface{}{}
	if cb.argLength == 0 {
		cb.funcVal.Call(cb.args)
		return
	}

	for _, boxedVal := range args {
		cb.args = append(cb.args, reflect.ValueOf(boxedVal))
	}

	if len(cb.args)%cb.argLength == 0 {
		resultVals := cb.funcVal.Call(cb.args)
		for _, resultVal := range resultVals {
			results = append(results, resultVal.Interface())
		}
		cb.args = []reflect.Value{}
	}
	return
}
