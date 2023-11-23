package usecases

import (
	"food-app/repository"
	"reflect"
)

type Recipe struct {
	Name        string
	Ingredients []Ingredient
	Version     int
}

type Ingredient struct {
	Name    string
	Serving string
	Type    FoodType
}

type Meat FoodType
type Fish FoodType
type Vegetable FoodType
type Legumes FoodType

type FoodType string

func RetrieveRecipes() []Recipe {
	return []Recipe{
		{
			Name: "Potato Galettes",
			Ingredients: []Ingredient{
				{
					Name:    "Potato",
					Serving: "4",
					Type:    "Vegetable",
				},
			},
			Version: 1,
		}}
}

func CreateRecipe(recipe Recipe) error {

	var sourceRecipe repository.Recipe
	mapStructFields(&recipe, &sourceRecipe)

	r, err := repository.StoreRecipeInRepository(sourceRecipe)

	if r == false {
		return err
	}
	return nil
}

func mapStructFields(src, dst interface{}) {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Ptr || dstValue.Kind() != reflect.Ptr {
		panic("Both source and destination must be pointers to structs")
	}

	srcValue = srcValue.Elem()
	dstValue = dstValue.Elem()

	if srcValue.Kind() != reflect.Struct || dstValue.Kind() != reflect.Struct {
		panic("Both source and destination must be pointers to structs")
	}

	srcType := srcValue.Type()
	dstType := dstValue.Type()

	// Iterate through fields of the source struct and map them to the destination struct
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		dstField, found := dstType.FieldByName(srcField.Name)

		// Only map fields with matching names and compatible types
		if found && srcField.Type == dstField.Type {
			dstFieldValue := dstValue.FieldByName(srcField.Name)
			dstFieldValue.Set(srcValue.Field(i))
		}
	}
}
