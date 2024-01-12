package postgre

import (
	"archi/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ElementRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewElementRepository(DB *gorm.DB, logger *zap.Logger) *ElementRepository {
	return &ElementRepository{DB: DB, Logger: logger}
}

func (r ElementRepository) GetByID(ctx context.Context, id string) ([]model.Element, error) {
	var elements []model.Element

	err := r.DB.Model(&model.Element{}).
		Select("elements.name, properties.name as name_prop, properties.value").
		Joins("JOIN properties ON elements.id = properties.parent_id").
		Where("elements.id = ?", id).
		Scan(&elements).Error

	if err != nil {
		r.Logger.Info("Error with Getting element by ID")
		return elements, err
	}
	fmt.Println(elements)
	return elements, nil
}
