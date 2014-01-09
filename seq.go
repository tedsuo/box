package box

import (
	"reflect"
)

type Box []interface{}

type Sequence chan Box

func NewSeq(input ...interface{}) (seq Sequence) {
	switch len(input) {
	case 0:
		return newEmptySeq()
	case 1:
		switch input := input[0].(type) {
		case Sequence:
			seq = input
		case Sequencer:
			seq = input.Seq()
		default:
			seq = newSeqFromValue(reflect.ValueOf(input))
		}
	default:
		seq = newSeqFromValue(reflect.ValueOf(input))
	}

	return
}

func newEmptySeq() Sequence {
	return make(Sequence)
}

func newSeqFromValue(inputVal reflect.Value) (seq Sequence) {
	seq = newEmptySeq()

	go func() {
		defer close(seq)
		switch inputVal.Kind() {
		case reflect.Map:
			seqNativeMapValue(seq, inputVal)
		case reflect.Slice:
			seqSliceValue(seq, inputVal)
		case reflect.Struct:
			seqStructValue(seq, inputVal)
		}
	}()

	return
}

func seqNativeMapValue(seq Sequence, inputVal reflect.Value) {
	for _, key := range inputVal.MapKeys() {
		val := inputVal.MapIndex(key)
		seq <- Box{key.Interface(),val.Interface()}
	}
	return
}

func seqSliceValue(seq Sequence, inputVal reflect.Value) {
	count := inputVal.Len()
	for i := 0; i < count; i++ {
		val := inputVal.Index(i)
		seq <- Box{val.Interface()}
	}
}

func seqStructValue(seq Sequence, inputVal reflect.Value) {
	inputType := inputVal.Type()
	FieldCount := inputVal.NumField()
	for i := 0; i < FieldCount; i++ {
		key := inputType.Field(i).Name
		val := inputVal.Field(i).Interface()
		seq <- Box{key,val}
	}
}
