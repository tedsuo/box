package box

import (
	"reflect"
)

type Sequencer interface {
	Seq() Sequence
}

type Box interface{}

type Sequence chan Box

func NewSeq(input ...interface{}) (seq Sequence) {
	if len(input) == 0 {
		return newEmptySeq()
	}

	switch input := input[0].(type) {
	case Sequence:
		seq = input
	case Sequencer:
		seq = input.Seq()
	default:
		seq = newSeqFromValue(reflect.ValueOf(input))
	}

	return
}

func newEmptySeq() Sequence {
	return make(Sequence)
}

func newSeqFromValue(inputVal reflect.Value) (seq Sequence) {
	switch inputVal.Kind() {
	case reflect.Map:
		seq = newSeqFromNativeMapValue(inputVal)
	case reflect.Slice:
		seq = newSeqFromSliceValue(inputVal)
	case reflect.Struct:
		seq = newSeqFromStructValue(inputVal)
	}
	return
}

func newSeqFromNativeMapValue(inputVal reflect.Value) (seq Sequence) {
	seq = newEmptySeq()

	go func() {
		defer close(seq)
		for _, key := range inputVal.MapKeys() {
			val := inputVal.MapIndex(key)
			seq <- key.Interface()
			seq <- val.Interface()
		}
	}()

	return
}

func newSeqFromSliceValue(inputVal reflect.Value) (seq Sequence) {
	seq = newEmptySeq()

	go func() {
		defer close(seq)
		count := inputVal.Len()
		for i := 0; i < count; i++ {
			val := inputVal.Index(i)
			seq <- val.Interface()
		}
	}()

	return
}

func newSeqFromStructValue(inputVal reflect.Value) (seq Sequence) {
	seq = newEmptySeq()
	inputType := inputVal.Type()
	FieldCount := inputVal.NumField()
	go func() {
		defer close(seq)
		for i := 0; i < FieldCount; i++ {
			key := inputType.Field(i).Name
			val := inputVal.Field(i).Interface()
			seq <- key
			seq <- val
		}
	}()
	return
}
