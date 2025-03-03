package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jalenarms1/foodgo/internal/db"
)

func HandlerGetFoodShopCategories(w http.ResponseWriter, r *http.Request) error {

	categories, err := db.GetFoodShopCategories()
	if err != nil {
		return err
	}

	fmt.Println(categories)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf(`{"categories": %v}`, categories))

	return nil
}

func HandleNewFoodShop(w http.ResponseWriter, r *http.Request) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var foodShop db.FoodShop
	err = json.Unmarshal(body, &foodShop)
	if err != nil {
		return err
	}

	fmt.Println(foodShop)

	err = foodShop.Insert()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
