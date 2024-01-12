package postgre

import (
	"archi/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewUserRepository(DB *gorm.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{DB: DB, Logger: logger}
}
func (r UserRepository) CreateUser(ctx context.Context, item model.User) (int, error) {
	fmt.Println("in repo user create")
	err := r.DB.Table("users").Create(&item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r UserRepository) GetUser(ctx context.Context, ID int) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (r UserRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (r UserRepository) Auth(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) DeleteUser(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.User{}, ID).Error
}
