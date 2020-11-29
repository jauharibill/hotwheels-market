package main

import (
	"github.com/go-chi/chi"
	"lelang-hotwheels/handler"
)

func router() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/", handler.Auth.Home)
	router.Post("/login", handler.Auth.Login)
	router.Post("/register", handler.Auth.Register)
	router.Post("/forgot", handler.Auth.ForgotPassword)

	return router
}
