package box

import ()

func Each(sequenceInput interface{}, callbacks ...interface{}) {
	cbStack := newCallbackStack(callbacks)
	for ß := range NewSeq(sequenceInput) {
		for _, cb := range cbStack {
			cb.Call(ß)
		}
	}
}

func Concat(args ...interface{}) (out Sequence) {
	out = NewSeq()
	go func() {
		defer close(out)
		for _, arg := range args {
			for ß := range NewSeq(arg) {
				out <- ß
			}
		}
	}()
	return
}

func Count(input interface{}) int {
	i := 0
	Each(input, func() {
		i++
	})
	return i
}

func Stream(sequenceInput interface{}, callbacks ...interface{}) (seq Sequence) {
	seq = NewSeq()
	cbStack := newCallbackStack(callbacks)
	go func() {
		defer close(seq)
		for boxedVal := range NewSeq(sequenceInput) {
			for _, cb := range cbStack {
				for ß := range NewSeq(cb.Call(boxedVal)) {
					seq <- ß
				}
			}
		}
	}()
	return
}
