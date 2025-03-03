package db

import (
	"context"
	"time"
)

type FoodShop struct {
	ID                  string             `json:"id"`
	AccountID           string             `json:"accountId"`
	Label               string             `json:"label"`
	Bio                 *string            `json:"bio"`
	FoodCategory        string             `json:"foodCategory"`
	Logo                string             `json:"logo"`
	IsDeliveryAvailable bool               `json:"isDeliveryAvailable"`
	Address             string             `json:"address"`
	City                string             `json:"city"`
	State               string             `json:"state"`
	ZipCode             string             `json:"zipCode"`
	Country             string             `json:"country"`
	Latitude            *float64           `json:"latitude"`
	Longitude           *float64           `json:"longitude"`
	MaxDeliveryRadius   *float64           `json:"maxDeliveryRadius"`
	DeliveryFee         *float64           `json:"deliveryFee"`
	CreatedAt           time.Time          `json:"createdAt"`
	FoodShopItems       []FoodShopItem     `json:"foodShopItems,omitempty"`
	FoodShopSchedule    []FoodShopSchedule `json:"foodShopSchedule,omitempty"`
}

func (f *FoodShop) Insert() error {
	_, err := Pool.Exec(context.Background(), "insert into FoodShop (Id, AccountId, Label, Bio, FoodCategory, Logo, IsDeliveryAvailable, Address, City, State, ZipCode, Country, Latitude, Longitude, MaxDeliveryRadius, DeliveryFee, CreatedAt) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)", f.ID, f.AccountID, f.Label, f.Bio, f.FoodCategory, f.Logo, f.IsDeliveryAvailable, f.Address, f.City, f.State, f.ZipCode, f.Country, f.Latitude, f.Longitude, f.MaxDeliveryRadius, f.DeliveryFee, f.CreatedAt)

	return err
}

func GetFoodShopCategories() (*[]string, error) {
	rows, err := Pool.Query(context.Background(), "select e.enumlabel from pg_enum e join pg_type t on t.oid = e.enumtypid where t.typname = 'foodcategory' order by e.enumlabel")
	if err != nil {
		return nil, err
	}

	var categories []string
	for rows.Next() {
		var cat string
		_ = rows.Scan(&cat)

		categories = append(categories, cat)
	}

	return &categories, nil
}
