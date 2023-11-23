package usecases_test

import (
	"food-app/usecases"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRecipes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Recipes", func() {

	It("should create a recipe", func() {
		recipe := usecases.Recipe{
			Name: "Test Recipe",
			Ingredients: []usecases.Ingredient{
				{
					Name:    "Potatoes",
					Serving: "4",
					Type:    "Vegetable",
				},
			},
			Version: 1,
		}
		err := usecases.CreateRecipe(recipe)
		Expect(err).To(BeNil())
	})

	It("should retrieve recipe", func() {
		r := usecases.RetrieveRecipes()
		Expect(r).To(HaveLen(1))
	})
})
