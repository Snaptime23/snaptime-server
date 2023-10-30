package dao

import (
	"context"
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func CreateUser(ctx context.Context, username, password string) (string, error) {
	user := &model.User{
		ID:       uuid.NewString(),
		UserName: username,
		Password: password,
	}
	err := DB.WithContext(ctx).Model(model.User{}).Create(user).Error
	if err != nil {
		log.Printf("Create User error: %v, user: %+v", err, user)
		return "", err
	}
	return user.ID, nil
}

func GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Model(model.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("Get User by Id error: %v", err)
		return &user, err
	}
	return &user, nil
}

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Model(model.User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, err
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, uid string, userMap *map[string]interface{}) (err error) {
	if err = DB.WithContext(ctx).Model(model.User{}).Where("id = ?", uid).Updates(&userMap).Error; err != nil {
		return err
	}
	return nil
}
