package box

type Sequencer interface {
	Seq() Sequence
}

type Collection interface {
	Sequencer
	Has(item interface{}) bool
	Delete(item interface{})
	Count() int
}
