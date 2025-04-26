package users

import "github.com/lcserny/go-authservice/src/config"

type Service struct {
	cfg *config.Config
	repo UserRepository
}

func NewUserService(cfg *config.Config, repo UserRepository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) GetUser(id string) string {
	return "a user"
}
