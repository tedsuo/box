package box

type Collection interface {
	Sequencer
	Has(item interface{}) bool
	Delete(item interface{})
	Count() int
}
