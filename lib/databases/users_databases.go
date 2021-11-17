package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

var user models.Users

var SelectUser = []string{
	user.Nama,
	user.Email,
}

func GetUser(id int) (interface{}, error) {
	if err := config.DB.Select(&SelectUser).First(&user, id).Error; err != nil {
		return nil, err
	}

	return SelectUser, nil
}

func CreateUser(user models.Users) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
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

func UpdateUser(id int, user models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}
