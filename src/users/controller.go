package users

import (
	"context"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/generated"
	"github.com/lcserny/go-authservice/src/logging"
)

var _ generated.GetSingleUserAPIServicer = (*Controller)(nil)

type Controller struct {
	cfg     *config.Config
	service *Service
}

func NewUsersController(cfg *config.Config, service *Service) *Controller {
	return &Controller{
		cfg:     cfg,
		service: service,
	}
}

func (c *Controller) GetUser(ctx context.Context, id string) (generated.ImplResponse, error) {
	user, err := c.service.GetUser(ctx, id)
	if err != nil {
		return generated.ImplResponse{}, err
	}
	return generated.Response(200, user), nil
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

	if resp.Body != nil {
		if err := generated.EncodeJSONResponse(resp.Body, &resp.Code, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (c *Controller) CreateUser(ctx context.Context, id string) (generated.ImplResponse, error) {
	user, err := c.service.CreateUser(ctx, id)
	if err != nil {
		return generated.ImplResponse{}, err
	}
	return generated.Response(200, user), nil
}

func (c *Controller) CreateUserAPI(w http.ResponseWriter, r *http.Request) {
	logging.Info(r.URL.Path + " " + r.Method)

	userID := strconv.Itoa(rand.Int())
	ctx := r.Context()

	resp, err := c.CreateUser(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.Body != nil {
		if err := generated.EncodeJSONResponse(resp.Body, &resp.Code, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
