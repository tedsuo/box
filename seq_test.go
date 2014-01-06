package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ƒ "github.com/tedsuo/box"
)

var _ = Describe("Seq", func() {
	Describe("New", func() {

		It("NewSeq() creates an empty seq", func() {
			seq := ƒ.NewSeq()
			i := 0

			go func() {
				defer close(seq)
				seq <- "foo"
			}()

			for val := range seq {
				Ω(val.(string)).Should(Equal("foo"))
				i++
			}

			Ω(i).Should(Equal(1))
		})

		It("NewSeq(map[string]string) creates a seq of key/val pairs", func() {
			input := map[string]string{
				"Name":    "Joe Bob",
				"Age":     "45",
				"Married": "true",
			}

			Ω(ƒ.Count(input)).Should(Equal(6))

			i := 0
			ƒ.Each(input, func(key, val string) {
				Ω(input[key]).Should(Equal(val))
				i++
			})
			Ω(i).Should(Equal(3))
		})

		It("NewSeq(struct{}) creates a seq of key/val pairs", func() {
			type Person struct {
				Name    string
				Age     int
				Married bool
			}

			joeBob := Person{
				Name:    "Joe Bob",
				Age:     45,
				Married: true,
			}

			Ω(ƒ.Count(joeBob)).Should(Equal(6))

			i := 0
			ƒ.Each(joeBob, func(key string, val interface{}) {
				switch key {
				case "Name":
					Ω(val.(string)).Should(Equal("Joe Bob"))
					i++
				case "Age":
					Ω(val.(int)).Should(Equal(45))
					i++
				case "Married":
					Ω(val.(bool)).Should(Equal(true))
					i++
				}
			})
			Ω(i).Should(Equal(3))
		})

	})
})
