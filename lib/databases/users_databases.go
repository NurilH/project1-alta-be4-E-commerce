package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

func GetUser(id int) (interface{}, error) {
	users := models.Users{}
	type get_user struct {
		Nama  string
		Email string
	}
	res := get_user{}
	err := config.DB.Model(users).Find(&res, id)
	rows_affected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return res, nil
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

func LoginUser(plan_pass string, user *models.Users) (interface{}, error) {
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	// log.Println(err)
	// log.Println("userpass", user.Password)
	if err != nil {
		return nil, err
	}
	// match, err := helper.CheckHashPassword(user.Password, plan_pass)
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plan_pass))
	// log.Println("match", match)
	if match != nil {
		return nil, match
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
