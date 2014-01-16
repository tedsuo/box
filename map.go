package box

import (
	"fmt"
)

// INTERFACE

type Map interface {
	Collection
	Get(key interface{}) interface{}
	Set(key, value interface{})
	Delete(item interface{})
	Keys() []interface{}
	IsNil(key interface{}) bool
	NotNil(key interface{}) bool
}

// INITIALIZATION

func NewMap(input ...interface{}) (mapp Map) {
	switch len(input) {
	case 0:
		return newEmptyMap()

	case 1:
		switch data := input[0].(type) {
		case Map:
			mapp = data
		case map[interface{}]interface{}:
			newMap := aMap(data)
			mapp = &newMap
		default:
			mapp = newEmptyMap()
			Each(data, mapp.Set)
		}

	default:
		panic("NewMap called with more than 1 argument")
	}

	return
}

func newEmptyMap() Map {
	return &aMap{}
}

// IMPLEMENTATION

type aMap map[interface{}]interface{}

func (data *aMap) Seq() (seq Sequence) {
	seq = NewSeq()
	go func() {
		defer close(seq)
		for key, val := range *data {
			seq <- Box{key, val}
		}
	}()
	return
}

func (data *aMap) String() string {
	return fmt.Sprintf("%+v", *data)
}

func (data *aMap) IsEmpty() bool {
	return data.Count() == 0
}

func (data *aMap) Count() int {
	return len(*data)
}

func (data *aMap) Has(key interface{}) bool {
	_, ok := (*data)[key]
	return ok
}

func (data *aMap) IsNil(key interface{}) bool {
	maybe, ok := (*data)[key]
	return ok && maybe == nil
}

func (data *aMap) NotNil(key interface{}) bool {
	maybe, ok := (*data)[key]
	return ok && maybe != nil
}

func (data *aMap) Keys() (keys []interface{}) {
	keys = make([]interface{}, 0, data.Count())

	for key := range *data {
		keys = append(keys, key)
	}

	return
}

func (data *aMap) Get(key interface{}) interface{} {
	return (*data)[key]
}

func (data *aMap) Set(key, value interface{}) {
	(*data)[key] = value
}

func (data *aMap) Delete(key interface{}) {
	delete(*data, key)
}
