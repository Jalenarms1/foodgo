package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
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
