package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
)

var user models.Users

func GetUser(id int) (interface{}, error) {
	users := models.Users{}
	type get_user struct {
		ID    uint
		Nama  string
		Email string
	}
	err := config.DB.Find(&users, id)
	rows_affected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_user{users.ID, users.Nama, users.Email}, nil
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

func LoginUser(user *models.Users) (interface{}, error) {
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
	if err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user.Token, nil
}
