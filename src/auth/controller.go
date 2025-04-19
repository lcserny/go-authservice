package auth

import (
	"github.com/lcserny/go-authservice/src/logging"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)
	w.Write([]byte(signIn("12")))
}
