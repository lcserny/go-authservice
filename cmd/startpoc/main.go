package main

import (
	"fmt"
	"net/http"

	"github.com/lcserny/go-authservice/pkg/mongodb"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/lcserny/go-authservice/src/web"
)

func main() {
	cfg := config.NewConfig()
	db := mongodb.NewMongoRepositoryProvider(cfg)
	router := web.NewRouter(cfg, db)
	logging.Info(fmt.Sprintf("Starting %s on port: %d", cfg.Application.Name, cfg.Application.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Application.Port), router)
	if err != nil {
		println(err.Error())
	}
}
