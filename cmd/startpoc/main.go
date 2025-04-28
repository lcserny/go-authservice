package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lcserny/go-authservice/pkg/mongodb"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/lcserny/go-authservice/src/web"
	slogchi "github.com/samber/slog-chi"
)

func main() {
	cfg := config.NewConfig()

	repoProvider := mongodb.NewMongoRepositoryProvider(cfg)

	r := web.InitRouter(cfg, repoProvider)
	// FIXME: cant add Use() middleware after controllers are added in router, provide a custom NewRouter()?
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

	logging.Info(fmt.Sprintf("Starting %s on port: %d", cfg.Application.Name, cfg.Application.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Application.Port), r)
	if err != nil {
		println(err.Error())
	}
}
