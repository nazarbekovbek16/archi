package storage

import (
	"archi/config"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	Pg      *gorm.DB
	User    IUserRepository
	Element IElementRepository
}

func NewStorage(logger *zap.Logger, ctx context.Context, cfg *config.Config) (*Storage, error) {
	var storage Storage

	//pgDB, err := postgre.OpenDB(cfg)
	//if err != nil {
	//	logger.Error("Dial error", zap.Error(err))
	//	return nil, err
	//}

	//storage.User = postgre.NewUserRepository(pgDB, logger)
	//storage.Element = postgre.NewElementRepository(pgDB, logger)

	return &storage, nil
}
