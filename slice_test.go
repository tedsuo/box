package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ƒ "github.com/tedsuo/box"
)

var _ = Describe("Slice", func() {
	Describe("New", func() {

		It("NewSlice() creates an empty slice", func() {
			s := ƒ.NewSlice()
			Ω(s.Count()).Should(Equal(0))
		})

		It("NewMap([]string) creates a slice of vals", func() {

			s := ƒ.NewSlice([]string{
				"Joe Bob",
				"45",
				"true",
			})

			Ω(s.Count()).Should(Equal(3))
			Ω(s.Get(0)).Should(Equal("Joe Bob"))
			Ω(s.Get(1)).Should(Equal("45"))
			Ω(s.Get(2)).Should(Equal("true"))
		})

		It("NewSlice(struct{}) creates a slice of struct values", func() {
			type Person struct {
				Name    string
				Age     int
				Married bool
			}

			mapp := ƒ.NewSlice(Person{
				Name:    "Joe Bob",
				Age:     45,
				Married: true,
			})

			Ω(mapp.Count()).Should(Equal(3))
			Ω(mapp.Get(0)).Should(Equal("Joe Bob"))
			Ω(mapp.Get(1)).Should(Equal(45))
			Ω(mapp.Get(2)).Should(Equal(true))
		})

		It("NewSlice(a,b,c) creates a slice of args", func() {
			mapp := ƒ.NewSlice("Joe Bob", 45, true)

			Ω(mapp.Count()).Should(Equal(3))
			Ω(mapp.Get(0)).Should(Equal("Joe Bob"))
			Ω(mapp.Get(1)).Should(Equal(45))
			Ω(mapp.Get(2)).Should(Equal(true))
		})
	})
})
