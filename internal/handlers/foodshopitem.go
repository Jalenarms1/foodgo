package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jalenarms1/foodgo/internal/db"
)

func HandleGetFoodItemCategories(w http.ResponseWriter, r *http.Request) error {

	categories, err := db.GetFoodItemCategories()
	if err != nil {
		return err
	}

	fmt.Println(categories)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf(`{"categories": %v}`, categories))

	return nil
}

func HandlerNewFoodShopItem(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var foodShopItem db.FoodShopItem
	err = json.Unmarshal(body, &foodShopItem)
	if err != nil {
		return err
	}

	fmt.Println(foodShopItem)

	err = foodShopItem.Insert()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
