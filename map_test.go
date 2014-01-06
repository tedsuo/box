package box_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ƒ "github.com/tedsuo/box"
)

var _ = Describe("Map", func() {
	Describe("New", func() {

		It("NewMap() creates an empty map", func() {
			mapp := ƒ.NewMap()
			Ω(mapp.Count()).Should(Equal(0))
		})

		It("NewMap(map[string]string) creates a map of key vals", func() {

			mapp := ƒ.NewMap(map[string]string{
				"Name":    "Joe Bob",
				"Age":     "45",
				"Married": "true",
			})

			Ω(mapp.Count()).Should(Equal(3))
			Ω(mapp.Get("Name")).Should(Equal("Joe Bob"))
			Ω(mapp.Get("Age")).Should(Equal("45"))
			Ω(mapp.Get("Married")).Should(Equal("true"))
		})

		It("NewMap(struct{}) creates a map of struct fields", func() {
			type Person struct {
				Name    string
				Age     int
				Married bool
			}

			mapp := ƒ.NewMap(Person{
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

	Describe("unmarshalling JSON", func() {
		var (
			mapp ƒ.Map
			err  error
		)

		BeforeEach(func() {

			exampleJSON := []byte(`{
			  "name": "Joe Bob",
			  "age": 42,
			  "gpa": 4.2,
			  "env": {
			      "baz": "boom",
			      "foo": "bar"
			  }
			}`)
			foo := new(map[string]interface{})
			err = json.Unmarshal(exampleJSON, foo)
			mapp = ƒ.NewMap(*foo)
		})

		It("creates the correct key/val pairs", func() {
			Ω(err).ShouldNot(HaveOccurred())
			Ω(mapp.Get("name")).Should(Equal("Joe Bob"))
			Ω(mapp.Get("age")).Should(BeEquivalentTo(42))
			Ω(mapp.Get("gpa")).Should(BeEquivalentTo(4.2))
		})
	})

})
