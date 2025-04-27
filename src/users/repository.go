package users

import "context"

type User struct {
	ID string `bson:"id"`
}

// TODO make a CrudRepository that does basic CRUD, the impl only needs to provide the filters to find by?
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
}
