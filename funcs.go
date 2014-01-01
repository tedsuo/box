package box

import "reflect"

func Each(sequenceInput interface{}, callbackInput interface{}) {
	seq := NewSeq(sequenceInput)
	cb := reflect.ValueOf(callbackInput)
	argLength := cb.Type().NumIn()
	count := 0
	args := []reflect.Value{}

	for boxedVal := range seq {
		if argLength == 0 {
			cb.Call(args)
			continue
		}
		count++
		args = append(args, reflect.ValueOf(boxedVal))
		if count == argLength {
			cb.Call(args)
			count = 0
			args = []reflect.Value{}
		}
	}
}

func Merge(args ...interface{}) (mergedMap Map) {
	mergedMap = NewMap()
	Each(args, func(seq interface{}) {
		Each(seq, mergedMap.Set)
	})

	return
}

func Count(input interface{}) int {
	i := 0
	Each(input, func() {
		i++
	})
	return i
}
