package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

// function database untuk membuat produk baru
func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product.UsersID, nil
}

// function database untuk menampilkan seluruh produk
func GetAllProduct() (interface{}, error) {
	products := []models.Product{}
	err := config.DB.Select([]string{"id", "nama", "harga", "kategori", "deskripsi"}).Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

// function database untuk menampilkan produk by id
func GetProductById(id int) (interface{}, error) {
	product := models.Product{}
	type get_product struct {
		ID        uint
		Nama      string
		Harga     int
		Kategori  string
		Deskripsi string
	}
	err := config.DB.Find(&product, id)
	rows_affected := config.DB.Find(&product, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_product{product.ID, product.Nama, product.Harga, product.Kategori, product.Deskripsi}, nil
}

// function bantuan untuk mendapatkan id user pada tabel produk
func GetIDUserProduct(id int) (uint, error) {
	var product models.Product
	err := config.DB.Find(&product, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return product.UsersID, nil
}

// function database untuk menghapus produk by id
func DeleteProduct(id int) (interface{}, error) {
	var product models.Product
	check_product := config.DB.Find(&product, id).RowsAffected

	err := config.DB.Delete(&product).Error
	if err != nil || check_product > 0 {
		return nil, err
	}
	return product.UsersID, nil
}

//  function database untuk memperbarui data produk by id
func UpdateProduct(id int, products *models.Product) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
