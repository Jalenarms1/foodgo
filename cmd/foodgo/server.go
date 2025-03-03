package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Jalenarms1/foodgo/internal/handlers"
	"github.com/Jalenarms1/foodgo/internal/types"
)

type Server struct {
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Run() error {

	mux := http.NewServeMux()

	s.RegisterRoutes(mux)

	logMux := loggingMiddleware(mux)

	userMux := userMiddleware(logMux)

	fmt.Printf("http://localhost%s", s.Addr)

	return http.ListenAndServe(s.Addr, userMux)
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {

	fs := http.FileServer(http.Dir("./client/dist"))

	mux.Handle("/", fs)

	mux.HandleFunc("GET /api/get-me", errorCatchHandlerFunc(handlers.HandleGetMe, false))
	mux.HandleFunc("POST /api/user-account", errorCatchHandlerFunc(handlers.HandleNewAccount, false))
	mux.HandleFunc("POST /api/logout", errorCatchHandlerFunc(handlers.HandleLogout, true))
	mux.HandleFunc("POST /api/login", errorCatchHandlerFunc(handlers.HandleLogin, false))

	mux.HandleFunc("GET /api/food-shop-categories", errorCatchHandlerFunc(handlers.HandlerGetFoodShopCategories, true))
	mux.HandleFunc("GET /api/food-item-categories", errorCatchHandlerFunc(handlers.HandleGetFoodItemCategories, true))
	mux.HandleFunc("POST /api/new-food-shop", errorCatchHandlerFunc(handlers.HandleNewFoodShop, true))
	mux.HandleFunc("POST /api/new-food-item", errorCatchHandlerFunc(handlers.HandlerNewFoodShopItem, true))
	mux.HandleFunc("POST /api/new-food-schedule", errorCatchHandlerFunc(handlers.HandlerNewFoodShopSchedule, true))
}

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func errorCatchHandlerFunc(fn ErrorHandlerFunc, isProtected bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isProtected {
			uid := r.Context().Value(types.AuthKey)
			if uid == nil {
				fmt.Println("not authorized to view this route")
				http.Error(w, errors.New("not authorized to view this route").Error(), http.StatusBadRequest)
				return
			}

		}

		if err := fn(w, r); err != nil {
			fmt.Println(err)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(fmt.Sprintf(`{"error": "%s"}`, err.Error()))
			return
		}
	}
}
