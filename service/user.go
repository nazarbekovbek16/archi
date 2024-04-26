package service

import (
	"archi/model"
	"archi/storage"
	"context"
)

type UserService struct {
	repo *storage.Storage
}
type IUserService interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, ID int) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, ID int) error
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) Create(ctx context.Context, user model.User) error {
	_, err := s.repo.User.CreateUser(ctx, user)
	return err
}

func (s UserService) GetByID(ctx context.Context, ID int) (model.User, error) {
	user, err := s.repo.User.GetUser(ctx, ID)
	return user, err
}

func (s UserService) Get(ctx context.Context, username string) (model.User, error) {
	user, err := s.repo.User.GetByEmail(ctx, username)
	return user, err
}

func (s UserService) Update(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Delete(ctx context.Context, ID int) error {
	return s.repo.User.DeleteUser(ctx, ID)
}
