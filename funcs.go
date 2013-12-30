package box

import "reflect"

func Each(seq Sequence, cb interface{}) {
	cbFunc := reflect.ValueOf(cb)
	for _, key := range seq.Keys() {
		cbFunc.Call([]reflect.Value{
			reflect.ValueOf(key),
			reflect.ValueOf(seq.Get(key)),
		})
	}
}

func Merge(collection, otherCollection Collection) Collection {
	mergedMap := NewMap()

	iterator := func(key, value interface{}) {
		mergedMap.Set(key, value)
	}

	Each(collection, iterator)
	Each(otherCollection, iterator)

	return mergedMap
}
