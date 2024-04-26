package service

import (
	"archi/storage"
	"fmt"
	"go.uber.org/zap"
)

type Service struct {
	storage *storage.Storage
	User    IUserService
	Element IElementService
}

func NewService(logger *zap.Logger, storage *storage.Storage) (*Service, error) {
	if storage == nil {
		logger.Error("Storage pointer is empty")
		return nil, fmt.Errorf("storage is empty")
	}

	var service Service

	service.User = NewUserService(storage)
	service.Element = NewElementService(storage)

	return &service, nil
}
