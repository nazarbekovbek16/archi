package service

import (
	"archi/model"
	"archi/storage"
	"context"
)

type ElementService struct {
	repo *storage.Storage
}

func NewElementService(repo *storage.Storage) *ElementService {
	return &ElementService{repo: repo}
}

type IElementService interface {
	GetByID(ctx context.Context, ID string) ([]model.Element, error)
}

func (s ElementService) GetByID(ctx context.Context, ID string) ([]model.Element, error) {
	return s.repo.Element.GetByID(ctx, ID)
}
