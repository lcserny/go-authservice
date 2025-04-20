package auth

import (
	"net/http"

	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/logging"
)

type authController struct {
	cfg *config.Config
}

func NewAuthController(cfg *config.Config) *authController {
	return &authController{
		cfg: cfg,
	}
}

func (c *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)
	w.Write([]byte(signIn("12")))
}
