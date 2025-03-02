package main

import (
	"fmt"
	"net/http"

	"github.com/Jalenarms1/foodgo/internal/account"
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

	fmt.Printf("http://localhost%s", s.Addr)

	return http.ListenAndServe(s.Addr, logMux)
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {

	fs := http.FileServer(http.Dir("./client/dist"))

	mux.Handle("/", fs)

	mux.HandleFunc("POST /user-account", errorHandlerFunc(account.HandleNewAccount))
}

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func errorHandlerFunc(fn ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
