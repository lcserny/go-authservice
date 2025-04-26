package users

import "context"

type User struct {
	ID   string
}

type UserRepository interface {
	 GetUserByID(ctx context.Context, id string) (*User, error)
}