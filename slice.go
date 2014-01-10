package box

type Slice interface {
	Collection
	At(index int) interface{}
	Set(key, value interface{})
	Push(item interface{})
	Pop() interface{}
}
