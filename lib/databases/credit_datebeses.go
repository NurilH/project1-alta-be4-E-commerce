package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateCredit(credit *models.Credit) (interface{}, error) {
	if err := config.DB.Create(&credit).Error; err != nil {
		return nil, err
	}
	return credit, nil
}

func DeleteCredit(id int) (interface{}, error) {
	var credit models.Credit
	check_credit := config.DB.Find(&credit, id).RowsAffected
	err := config.DB.Delete(&credit).Error
	if err != nil || check_credit > 0 {
		return nil, err
	}
	return credit.UsersID, nil
}

func GetIDUserCredit(id int) (uint, error) {
	var credit models.Credit
	err := config.DB.Find(&credit, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return credit.UsersID, nil
}
