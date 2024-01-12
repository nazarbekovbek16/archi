package storage

import (
	"archi/model"
	"context"
)

type IUserRepository interface {
	GetUser(ctx context.Context, ID int) (model.User, error)
	GetByEmail(ctx context.Context, username string) (model.User, error)
	Auth(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, ID int) error
	CreateUser(ctx context.Context, item model.User) (int, error)
}

type IElementRepository interface {
	GetByID(ctx context.Context, id string) ([]model.Element, error)
}
