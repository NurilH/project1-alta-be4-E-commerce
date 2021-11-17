package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func GetProductById(id int) (interface{}, error) {
	product := models.Product{}
	type get_product struct {
		ID        uint
		Nama      string
		Harga     int
		Deskripsi string
	}
	err := config.DB.Find(&product, id)
	rows_affected := config.DB.Find(&product, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_product{product.ID, product.Nama, product.Harga, product.Deskripsi}, nil
}
