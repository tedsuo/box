package f

type Collection interface {
	Sequence
	IsEmpty() bool
	Has(key interface{}) bool
	Set(key, value interface{})
	Delete(key interface{})
	Count() int
}

type Sequence interface {
	Keys() []interface{}
	Get(key interface{}) interface{}
}

type IteratorFunc func(key, val interface{})
