package repository

import (
	"encoding/json"
	"fmt"
	"os"
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

type FoodType string

func StoreRecipeInRepository(recipe Recipe) (bool, error) {
	dirPath := "."
	fileName := "database.json"

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return false, err
	}

	file, err := os.OpenFile(dirPath+"/"+fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false, err
	}
	defer file.Close()

	data, err := json.Marshal(recipe)

	if err != nil {
		fmt.Println("Error serializing data", err)
		return false, err
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return false, err
	}

	return true, nil
}
