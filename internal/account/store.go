package account

import (
	"context"

	"github.com/Jalenarms1/foodgo/internal/db"
)

type Account struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Account) Insert() error {

	_, err := db.Pool.Exec(context.Background(), `insert into "User" (Id, Email, Password) values ($1, $2, $3)`, a.Id, a.Email, a.Password)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountById(id string) (*Account, error) {

	var acct Account
	row := db.Pool.QueryRow(context.Background(), `select Id, Email from "User" where Id = $1`, id)

	err := row.Scan(&acct.Id, &acct.Email)
	if err != nil {
		return nil, err
	}

	return &acct, nil
}

func GetAccountByEmail(email string) (*Account, error) {

	var acct Account
	row := db.Pool.QueryRow(context.Background(), `select Id, Email from "User" where Email = $1`, email)

	err := row.Scan(&acct.Id, &acct.Email)
	if err != nil {
		return nil, err
	}

	return &acct, nil
}
