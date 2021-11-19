package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project_altabe4_1/config"
	"project_altabe4_1/constants"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type ProductResponse struct {
	Message string
	Data    models.Product
}

var (
	mock_data_product = models.Product{
		Nama:      "andri",
		Harga:     99000,
		Kategori:  "kategori",
		Deskripsi: "deskripsi",
		UsersID:   1,
	}
	mock_data_user = models.Users{
		Email:    "andri@gmail.com",
		Password: "bismillah",
	}
	mock_data_login = models.Users{
		Email:    "andri@gmail.com",
		Password: "bismillah",
	}
)

func InitEchoProducts() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

func InsertUser() error {
	if err := config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

func InsertProduct() error {
	if err := config.DB.Save(&mock_data_product).Error; err != nil {
		return err
	}
	return nil
}

//test create product
func TestCreateProductController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/products",
		expectCode: http.StatusOK,
	}

	e := InitEchoProducts()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	body, err := json.Marshal(mock_data_product)
	if err != nil {
		t.Error(t, err, "error")
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductControllersTesting())(c)

	bodyrecponses := rec.Body.String()
	var product ProductResponse
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test create product error
func TestCreateProductControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "jwt/products",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	body, err := json.Marshal(mock_data_product)
	if err != nil {
		t.Error(t, err, "error")
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductControllersTesting())(c)

	bodyrecponses := rec.Body.String()
	var product ProductResponse
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test get all product
func TestGetAllProductController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/products/",
		expectCode: http.StatusOK,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	middleware.JWT([]byte(constants.SECRET_JWT))(GetProductsControllerTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test get all product error
func TestGetAllProductControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "jwt/products/",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(GetProductsControllerTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test get product by id
func TestGetProductByIDController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/products/:id",
		expectCode: http.StatusOK,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetProductByIdControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test get product by id error
func TestGetProductByIDControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	config.DB.Migrator().DropTable(models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(GetProductByIdControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test delete product by id
func TestDeleteProductByIDController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/products/:id",
		expectCode: http.StatusOK,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test delete product by id error
func TestDeleteProductByIDControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Access Forbidden",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	config.DB.Migrator().DropTable(models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test update product by id
func TestUpdateProductByIDController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/products/:id",
		expectCode: http.StatusOK,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test update product by id error
func TestUpdateProductByIDControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Access Forbidden",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("1")
	config.DB.Migrator().DropTable(models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test get product by id false param
func TestGetProductByIDControllerFalseParam(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "False Param",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("x")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetProductByIdControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test delete product by id false param
func TestDeleteProductByIDControllerFalseParam(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "False Param",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("x")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}

//test update product by id false param
func TestUpdateProductByIDControllerFalseParam(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "False Param",
		path:       "jwt/products/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEchoProducts()
	InsertUser()
	InsertProduct()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(testCases.path)
	c.SetParamNames("id")
	c.SetParamValues("x")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllersTesting())(c)

	var product ProductResponse
	bodyrecponses := rec.Body.String()
	err = json.Unmarshal([]byte(bodyrecponses), &product)
	if err != nil {
		assert.Error(t, err, "error")
	}

	assert.Equal(t, testCases.expectCode, rec.Code)
	assert.Equal(t, testCases.name, product.Message)

}
