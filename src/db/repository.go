package db

import (
	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/users"
)

type RepositoryProvider interface {
	GetUserRepository() users.UserRepository
	GetAuthRepository() auth.AuthRepository
}