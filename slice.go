package box

import (
	"fmt"
)

// INTERFACE

type Slice interface {
	Collection
	Get(index interface{}) interface{}
	Set(index, value interface{})
	Push(item interface{})
	Pop() interface{}
	Insert(index, value interface{})
}

// INITIALIZATION

func NewSlice(input ...interface{}) (s Slice) {
	switch len(input) {
	case 0:
		return newEmptySlice()

	case 1:
		switch data := input[0].(type) {
		case Slice:
			s = data
		case []interface{}:
			newSlice := aSlice(data)
			s = &newSlice
		default:
			s = newEmptySlice()
			Each(data, s.Push)
		}

	default:
		s = newEmptySlice()
		Each(input, s.Push)
	}

	return
}

func newEmptySlice() Slice {
	return &aSlice{}
}

// IMPLEMENTATION

type aSlice []interface{}

func (data *aSlice) Seq() (seq Sequence) {
	seq = NewSeq()
	go func() {
		defer close(seq)
		for i, val := range *data {
			seq <- Box{i, val}
		}
	}()
	return
}

func (data *aSlice) String() string {
	return fmt.Sprintf("%+v", *data)
}

func (data *aSlice) IsEmpty() bool {
	return data.Count() == 0
}

func (data *aSlice) Count() int {
	return len(*data)
}

func (data *aSlice) Has(item interface{}) bool {
	for _, val := range *data {
		if Equals(val, item) {
			return true
		}
	}

	return false
}

func (data *aSlice) Get(index interface{}) interface{} {
	return (*data)[Int(index)]
}

func (data *aSlice) Set(index, value interface{}) {
	(*data)[Int(index)] = value
}

func (data *aSlice) Push(item interface{}) {
	*data = append(*data, item)
}

func (data *aSlice) Pop() interface{} {
	lastIndex := data.Count() - 1
	value := (*data)[lastIndex]
	*data = (*data)[:lastIndex]
	return value
}

func (data *aSlice) Insert(index, value interface{}) {
	i := Int(index)
	*data = append(*data, nil)
	copy((*data)[i+1:], (*data)[i:])
	(*data)[i] = value
}
