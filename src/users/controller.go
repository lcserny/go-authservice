package users

import (
	"net/http"

	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/logging"
)

type usersController struct {
	cfg *config.Config
}

func NewUsersController(cfg *config.Config) *usersController {
	return &usersController{
		cfg: cfg,
	}
}

func (c *usersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)
	w.Write([]byte(getUser("1")))
}
