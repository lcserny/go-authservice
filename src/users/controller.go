package users

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/generated"
	"github.com/lcserny/go-authservice/src/logging"
)

var _ generated.GetSingleUserAPIServicer = (*usersController)(nil)

type usersController struct {
	cfg *config.Config
}

func NewUsersController(cfg *config.Config) *usersController {
	return &usersController{
		cfg: cfg,
	}
}

// TODO: do actual work, call service and transform to ImplResponse here
func (c *usersController) GetUser(context.Context, string) (generated.ImplResponse, error) {
	return generated.ImplResponse{}, nil
}

func (c *usersController) GetUserAPI(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	resp, err, := c.GetUser(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.Code)
	if resp.Body != nil {
		w.Write(resp.Body)
	}
}
