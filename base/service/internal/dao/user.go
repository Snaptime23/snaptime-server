package dao

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/google/uuid"
	"log"
)

func CreateUser(ctx context.Context, username, password, avtar, des, email string) (string, error) {
	user := &model.User{
		UserID:   uuid.NewString(),
		UserName: username,
		Password: password,
		Avatar:   avtar,
		Bio:      des,
		Email:    email,
	}
	err := DB.WithContext(ctx).Model(model.User{}).Create(user).Error
	if err != nil {
		log.Printf("Create User error: %v, user: %+v", err, user)
		return "", err
	}
	return user.UserID, nil
}

func GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	err := DB.WithContext(ctx).Model(model.User{}).Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		log.Printf("Get User by Id error: %v", err)
		return &user, err
	}
	return &user, nil
}

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{}
	err := DB.WithContext(ctx).Model(model.User{}).Where("user_name = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(ctx context.Context, uid string, userMap *map[string]interface{}) (err error) {
	if err = DB.WithContext(ctx).Model(model.User{}).Where("user_id = ?", uid).Updates(&userMap).Error; err != nil {
		return err
	}
	return nil
}
