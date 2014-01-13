package box

// INTERFACE

type Slice interface {
	Collection
	At(index int) interface{}
	Set(key, value interface{})
	Push(item interface{})
	Pop() interface{}
}

// INITIALIZATION

func NewSlice(input ...interface{}) (slicee Slice) {
	if len(input) == 0 {
		return newEmptySlice()
	}

	switch data := input[0].(type) {
	case Slice:
		slicee = data
	case []interface{}:
		newSlice := aSlice(data)
		slicee = &newSlice
	default:
		slicee = newEmptySlice()
		Each(data, slicee.Push)
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

func (data *aSlice) Has(key interface{}) bool {
	_, ok := (*data)[key]
	return ok
}

func (data *aSlice) At(index int) interface{} {
	return (*data)[key]
}

func (data *aSlice) Push(key interface{}) interface{} {
	return (*data)[key]
}

func (data *aSlice) Pop() interface{} {
	(*data)[key] = value
}
