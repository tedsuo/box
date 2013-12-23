package f

type OnEach func(key, val interface{})

func Each(seq Sequence, cb OnEach) {
	for _, key := range seq.Keys() {
		cb(key, seq.Get(key))
	}
}

func Merge(collection, otherCollection Collection) Collection {
	mergedMap := NewEmptyMap()

	iterator := func(key, value interface{}) {
		mergedMap.Set(key, value)
	}

	Each(collection, iterator)
	Each(otherCollection, iterator)

	return mergedMap
}
