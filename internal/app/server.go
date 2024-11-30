package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aidosgal/gust/internal/http/handler"
	"github.com/aidosgal/gust/internal/repository"
	"github.com/aidosgal/gust/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)

	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Post("/login", userHandler.HandleLogin)
			router.Post("/register", userHandler.HandleRegister)
		})
	})

	log.Print("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
