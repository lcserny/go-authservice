package users

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/generated"
	"github.com/lcserny/go-authservice/src/logging"
)

var _ generated.GetSingleUserAPIServicer = (*Controller)(nil)

type Controller struct {
	cfg *config.Config
}

func NewUsersController(cfg *config.Config) *Controller {
	return &Controller{
		cfg: cfg,
	}
}

// TODO: do actual work, call service and transform to ImplResponse here
func (c *Controller) GetUser(context.Context, string) (generated.ImplResponse, error) {
	return generated.ImplResponse{200, "body"}, nil
}

func (c *Controller) GetUserAPI(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	resp, err := c.GetUser(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.Code)
	if resp.Body != nil {
		w.Write([]byte(resp.Body.(string)))
	}
}
