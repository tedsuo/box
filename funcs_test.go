package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/box"
)

var _ = Describe("Funcs", func() {

	Describe("Each: Map", func() {
		var coll box.Map
		var keys = []string{"a", "b", "c"}
		var values = []string{"apple", "berry", "cantelope"}

		BeforeEach(func() {
			coll = box.NewMap()
			for i, key := range keys {
				coll.Set(key, values[i])
			}
		})

		It("iterates over untyped keys and values", func() {
			count := 0
			box.Each(coll, func(key, val interface{}) {
				Ω(key.(string)).Should(Equal(keys[count]))
				Ω(val.(string)).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

		It("iterates over untyped keys and values", func() {
			count := 0
			box.Each(coll, func(key, val string) {
				Ω(key).Should(Equal(keys[count]))
				Ω(val).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

	})
})
