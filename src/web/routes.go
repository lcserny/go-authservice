package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/users"
)

func UserRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", users.GetUsers)
	return r
}

func AuthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", auth.SignIn)
	return r
}
