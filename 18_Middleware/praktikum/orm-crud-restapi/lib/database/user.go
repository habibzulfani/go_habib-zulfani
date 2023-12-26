package database

import (
	"orm-crud-restapi/config"
	"orm-crud-restapi/middlewares"
	"orm-crud-restapi/models"
)

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	e := config.DB.Find(&user, userId).Error
	if e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error

	if err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
