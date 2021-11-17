package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

var user models.Users

type get_user struct {
	Nama  string
	Email string
}

func GetUser(id int) (interface{}, error) {
	err := config.DB.Find(&user, id).Error
	rowsAffected := config.DB.Find(&user, id).RowsAffected
	if err != nil || rowsAffected < 1 {
		return nil, err
	}
	return get_user{user.Nama, user.Email}, nil
}

func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id int, user *models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}
