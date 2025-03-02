package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Jalenarms1/foodgo/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		if strings.HasSuffix(origin, os.Getenv("CLIENT_DOMAIN")) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		cookies := r.Cookies()

		fmt.Println(cookies)

		slog.Info("New Request: ", "Path", r.URL.Path, "Host", r.Host, "Addr", r.RemoteAddr, "Time", time.Now().Unix())

		next.ServeHTTP(w, r)
	})
}

func userMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(string(types.AuthKey))
		if err != nil || cookie == nil {
			fmt.Println(err)

		}
		var ctx context.Context
		if cookie != nil {
			token := cookie.Value
			fmt.Println(token)
			claims := &types.Claims{}
			jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(jwtToken.Valid)
			fmt.Println(claims.Uid)
			if claims.Uid != "" && jwtToken.Valid {
				fmt.Println(claims)

				ctx = context.WithValue(context.Background(), types.AuthKey, claims.Uid)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)

	})
}
