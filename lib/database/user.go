package database

import (
	"project/config"
	"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(id int, user interface{}) (interface{}, error) {
	if err := config.DB.Find(&user, "id=?", id).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
