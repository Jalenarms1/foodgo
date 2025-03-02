package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(acctInfo.Password), 10)
	if err != nil {
		return err
	}

	fmt.Println(acctInfo)
	fmt.Println(string(hashedPass))

	token, err := generateJWT(acctInfo.Email)
	if err != nil {
		return err
	}

	fmt.Println(token)
	// isDev := os.Getenv("IS_DEV")

	cookie := &http.Cookie{
		Name:     "foodgo-auth",
		Value:    token,
		Path:     "/",
		MaxAge:   3600 * 24,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)

	return nil
}

func generateJWT(email string) (string, error) {
	signingKey := os.Getenv("JWT_SECRET")

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(((24 * 365) * time.Hour)).Unix(),
	})

	token, err := claim.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil

}
