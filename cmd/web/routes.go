package main

import (
	"github.com/bmizerany/pat"
	"go-learn-web/pkg/config"
	"go-learn-web/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
