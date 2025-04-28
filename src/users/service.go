package users

import (
	"context"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/generated"
)

type Service struct {
	cfg  *config.Config
	repo UserRepository
}

func NewService(cfg *config.Config, repo UserRepository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) GetUser(ctx context.Context, userId string) (generated.ImplResponse, error) {
	panic("implement me")
}

func (s *Service) GetUsers(ctx context.Context, page int32, limit int32, username string, firstName string, lastName string) (generated.ImplResponse, error) {
	panic("implement me")
}

func (s *Service) Register(ctx context.Context, username string, password string, firstName string, lastName string) (generated.ImplResponse, error) {
	panic("implement me")
}

func (s *Service) UpdateUser(ctx context.Context, userId string, userData generated.UserData) (generated.ImplResponse, error) {
	panic("implement me")
}

//func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
//	user, err := s.repo.GetUserByID(ctx, id)
//	if err != nil {
//		return nil, errors.New("Error getting user: " + err.Error())
//	}
//	return user, nil
//}
//
//func (s *Service) CreateUser(ctx context.Context, id string) (*User, error) {
//	user := User{ID: id}
//	if err := s.repo.CreateUser(ctx, &user); err != nil {
//		return nil, errors.New("Error creating user: " + err.Error())
//	}
//	return &user, nil
//}
