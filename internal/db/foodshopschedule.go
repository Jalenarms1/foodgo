package db

import "context"

type FoodShopSchedule struct {
	Id         string `json:"id"`
	FoodShopID string `json:"foodShopId"`
	OpenDt     string `json:"openDt"`
	CloseDt    string `json:"closeDt"` // -- 20250101
	OpenTm     string `json:"openTm"`  // -- 0930, 1330
	CloseTm    string `json:"closeTm"` //
	CreatedAt  string `json:"createdAt"`
}

func (fi *FoodShopSchedule) Insert() error {
	query := `
		INSERT INTO FoodShopSchedule (
			Id, FoodShopId, OpenDt, CloseDt, OpenTm, CloseTm, CreatedAt
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := Pool.Exec(context.Background(), query, fi.Id, fi.FoodShopID, fi.OpenDt, fi.CloseDt, fi.OpenTm, fi.CloseTm, fi.CreatedAt)

	return err
}
