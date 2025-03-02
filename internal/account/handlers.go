package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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

	var acctInfo Account
	err = json.Unmarshal(body, &acctInfo)
	if err != nil {
		return err
	}

	if acctInfo.Email == "" || acctInfo.Password == "" {
		return errors.New("provide both an email and password")
	}

	existingAcct, _ := GetAccountByEmail(acctInfo.Email)
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

func HandleGetMe(w http.ResponseWriter, r *http.Request) error {

	uid := r.Context().Value(types.AuthKey).(string)

	fmt.Println(uid)

	acct, err := GetAccountById(uid)
	if err != nil {
		return err
	}

	fmt.Println(acct)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))

	return nil
}
