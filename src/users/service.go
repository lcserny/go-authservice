package users

import (
	"context"
	"errors"
	"github.com/lcserny/go-authservice/src/config"
)

type Service struct {
	cfg  *config.Config
	repo UserRepository
}

func NewUserService(cfg *config.Config, repo UserRepository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.New("Error getting user: " + err.Error())
	}
	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, id string) (*User, error) {
	user := User{ID: id}
	if err := s.repo.CreateUser(ctx, &user); err != nil {
		return nil, errors.New("Error creating user: " + err.Error())
	}
	return &user, nil
}
