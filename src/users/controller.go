package users

import (
	"github.com/lcserny/go-authservice/src/logging"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)
	w.Write([]byte(getUser("1")))
}
