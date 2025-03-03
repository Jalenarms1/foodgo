package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Jalenarms1/foodgo/internal/db"
	"github.com/Jalenarms1/foodgo/internal/types"
	"github.com/Jalenarms1/foodgo/internal/utils"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleNewAccount(w http.ResponseWriter, r *http.Request) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	fmt.Println(string(body))

	var acctInfo db.Account
	err = json.Unmarshal(body, &acctInfo)
	if err != nil {
		return err
	}

	if acctInfo.Email == "" || acctInfo.Password == "" {
		return errors.New("provide both an email and password")
	}

	existingAcct, _ := db.GetAccountByEmail(acctInfo.Email)
	if existingAcct != nil {
		return errors.New("account with this email already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(acctInfo.Password), 10)
	if err != nil {
		return err
	}

	acctInfo.Password = string(hashedPass)

	fmt.Println(acctInfo)
	fmt.Println(string(hashedPass))

	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	acctInfo.Id = uid.String()

	token, err := utils.GenerateJWT(uid.String())
	if err != nil {
		return err
	}

	fmt.Println(token)
	isDev := os.Getenv("IS_DEV") != "true"

	cookie := &http.Cookie{
		Name:     "foodgo-auth",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(3600 * time.Hour),
		Secure:   isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)

	err = acctInfo.Insert()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var acctInfo *db.Account

	err = json.Unmarshal(body, &acctInfo)
	if err != nil {
		return err
	}

	existingUser, _ := db.GetUserWPasswordByEmail(acctInfo.Email)
	if existingUser == nil {
		return errors.New("existing user not found with the email provided")
	}

	isPasswordMatch := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(acctInfo.Password))

	if isPasswordMatch != nil {
		return errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(existingUser.Id)
	if err != nil {
		return err
	}

	fmt.Println(token)
	isDev := os.Getenv("IS_DEV") != "true"

	cookie := &http.Cookie{
		Name:     "foodgo-auth",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(3600 * time.Hour),
		Secure:   isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)

	return nil
}

func HandleGetMe(w http.ResponseWriter, r *http.Request) error {

	uid := r.Context().Value(types.AuthKey)
	if uid == nil {
		return errors.New("no authentication")
	}

	fmt.Println(uid)

	acct, err := db.GetAccountById(uid.(string))
	if err != nil {
		return err
	}

	fmt.Println(acct)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acct)

	return nil
}

func HandleLogout(w http.ResponseWriter, r *http.Request) error {

	http.SetCookie(w, &http.Cookie{
		Name:     string(types.AuthKey),
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Secure:   os.Getenv("IS_DEV") != "true",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	w.WriteHeader(http.StatusOK)

	return nil
}
