package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/users"
)

func UserRoutes(cfg *config.Config) *chi.Mux {
	ctrl := users.NewUsersController(cfg)
	r := chi.NewRouter()
	r.Get("/", ctrl.GetUsers)
	return r
}

func AuthRoutes(cfg *config.Config) *chi.Mux {
	ctrl := auth.NewAuthController(cfg)
	r := chi.NewRouter()
	r.Get("/", ctrl.SignIn)
	return r
}
