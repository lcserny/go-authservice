package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/lcserny/go-authservice/src/web"
	slogchi "github.com/samber/slog-chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(slogchi.New(logging.Logger()))
	r.Use(middleware.Recoverer)
	r.Mount("/users", web.UserRoutes())
	r.Mount("/auth", web.AuthRoutes())
	http.ListenAndServe(":3000", r)
}
