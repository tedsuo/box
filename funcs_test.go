package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ß "github.com/tedsuo/box"
)

var _ = Describe("Funcs", func() {

	Describe("Each: Map", func() {
		var coll ß.Map
		var keys = []string{"a", "b", "c"}
		var values = []string{"apple", "berry", "cantelope"}

		BeforeEach(func() {
			coll = ß.NewMap()
			for i, key := range keys {
				coll.Set(key, values[i])
			}
		})

		It("iterates over untyped keys and values", func() {
			count := 0
			ß.Each(coll, func(key, val interface{}) {
				Ω(key.(string)).Should(Equal(keys[count]))
				Ω(val.(string)).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

		It("iterates over typed keys and values", func() {
			count := 0
			ß.Each(coll, func(key, val string) {
				Ω(key).Should(Equal(keys[count]))
				Ω(val).Should(Equal(values[count]))
				count++
			})
			Ω(count).Should(Equal(len(keys)))
		})

	})

	Describe("Merge: Map", func() {
		var mapA, mapB, mapC ß.Map
		BeforeEach(func() {
			mapA = ß.NewMap(map[string]string{
				"a": "apple",
				"b": "banana",
			})
			mapB = ß.NewMap(map[string]string{
				"c": "canberry",
				"d": "dandelion",
			})

			mapC = ß.NewMap(map[string]string{
				"e": "eclair",
				"f": "frappe",
			})
		})
		It("merges three maps into one", func() {
			mapp := ß.Merge(mapA, mapB, mapC)
			Ω(mapp.Count()).Should(Equal(6))
		})
	})
})
