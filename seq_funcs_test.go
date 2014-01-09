package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ƒ "github.com/tedsuo/box"
)

var _ = Describe("Funcs", func() {

	Describe("Each: Map", func() {
		var coll ƒ.Map
		var keys = []string{"a", "b", "c"}
		var values = []string{"apple", "berry", "cantelope"}

		BeforeEach(func() {
			coll = ƒ.NewMap()
			for i, key := range keys {
				coll.Set(key, values[i])
			}
		})

		It("iterates over untyped keys and values", func() {
			count := 0
			ƒ.Each(coll, func(key, val interface{}) {
				Ω(key.(string)).Should(Equal(keys[count]))
				Ω(val.(string)).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

		It("iterates over typed keys and values", func() {
			count := 0
			ƒ.Each(coll, func(key, val string) {
				Ω(key).Should(Equal(keys[count]))
				Ω(val).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

		It("iterates over only values", func() {
			count := 0
			ƒ.Each(coll, func(val string) {
				Ω(val).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(values)))
		})

	})

	Describe("Concat: Map", func() {
		var (
			mapA, mapB, mapC ƒ.Map
			seq              ƒ.Sequence
		)
		BeforeEach(func() {
			mapA = ƒ.NewMap(map[string]string{
				"a": "apple",
				"b": "banana",
			})
			mapB = ƒ.NewMap(map[string]string{
				"c": "canberry",
				"d": "dandelion",
			})
			mapC = ƒ.NewMap(map[string]string{
				"e": "eclair",
				"f": "frappe",
			})

			seq = ƒ.Concat(mapA, mapB, mapC)
		})

		It("merges three maps into one", func() {
			mapp := ƒ.NewMap(seq)
			Ω(mapp.Count()).Should(Equal(6))
		})
	})
})
