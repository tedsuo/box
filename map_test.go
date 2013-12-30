package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/box"
)

var _ = Describe("Map", func() {
	Describe("New", func() {

		It("NewMap() creates an empty map", func() {
			mapp := box.NewMap()
			Ω(mapp.Count()).Should(Equal(0))
		})

		It("NewMap(struct{}) creates a map of struct fields", func() {
			type Person struct {
				Name    string
				Age     int
				Married bool
			}

			mapp := box.NewMap(Person{
				Name:    "Joe Bob",
				Age:     45,
				Married: true,
			})

			Ω(mapp.Count()).Should(Equal(3))
			Ω(mapp.Get("Name")).Should(Equal("Joe Bob"))
			Ω(mapp.Get("Age")).Should(Equal(45))
			Ω(mapp.Get("Married")).Should(Equal(true))
		})

	})
})
