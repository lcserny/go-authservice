package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/lcserny/go-authservice/src/web"
	slogchi "github.com/samber/slog-chi"
)

func main() {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	r.Use(slogchi.New(logging.Logger()))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	r.Route(cfg.ContextPath, func(r chi.Router) {
		r.Mount("/users", web.UserRoutes(cfg))
		r.Mount("/auth", web.AuthRoutes(cfg))
	})

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), r)
	if err != nil {
		println(err.Error())
	}
}
