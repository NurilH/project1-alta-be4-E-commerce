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

type UserResponse struct {
	Message string
	Data    models.Users
}

type Login struct {
	Email    string
	Password string
}

var (
	mock_data_user = models.Users{
		Nama:     "andri",
		Email:    "andri@gmail.com",
		Password: "bismillah",
	}
	mock_data_login = models.Users{
		Email:    "andri@gmail.com",
		Password: "bismillah",
	}
)

func InitEcho() *echo.Echo {
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

//test get user
func TestGetUserControllers(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Successful Operation",
		path: "jwt/users/:id",
		code: http.StatusOK,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
		assert.Equal(t, "andri", user.Data.Nama)
	})

}

//test get user error
func TestGetUserControllersError(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{
		name: "Bad Request",
		path: "/users/:id",
		code: http.StatusBadRequest,
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues("1")
	context.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Users{})

	if assert.NoError(t, GetUserControllers(context)) {

		var user UserResponse
		rec_body := rec.Body.String()
		err := json.Unmarshal([]byte(rec_body), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)

	}
}

//test create user
func TestCreateUserController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "/users",
		expectCode: http.StatusOK,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}

}

//test create user error
func TestCreateUserControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "/users",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	config.DB.Migrator().DropTable(models.Users{})

	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	//send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}

}

//test update user
func TestUpdateUserController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "jwt/users/:id",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
		// assert.Equal(t, "andri", user.Data.Nama)
	})
}

//test update user error
func TestUpdateUserControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "jwt/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	config.DB.Migrator().DropTable(models.Users{})

	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("PUT /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
	})

}

func TestUpdateUserControllerForbidden(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Access Forbidden",
		path:       "jwt/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("100")
	config.DB.Migrator().DropTable(models.Users{})

	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("PUT /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
	})

}

//test delete user
func TestDeleteUserController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "/users/:id",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

//test delete user error
func TestDeleteUserControllerError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "jwt/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	config.DB.Migrator().DropTable(models.Users{})
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

func TestDeleteUserControllerForbidden(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Access Forbidden",
		path:       "jwt/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("100")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("DELETE /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

// test get user error param
func TestGetUserControllersFalseParam(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{
		name: "False Param",
		path: "/users/:id",
		code: http.StatusBadRequest,
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues("x")
	context.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Users{})

	if assert.NoError(t, GetUserControllers(context)) {

		var user UserResponse
		rec_body := rec.Body.String()
		err := json.Unmarshal([]byte(rec_body), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)

	}
}

// test update user error param
func TestUpdateUserControllersFalseParam(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "False Param",
		path:       "jwt/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("x")
	config.DB.Migrator().DropTable(models.Users{})

	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("PUT /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})
}

// test delete user error param
func TestDeleteUserControllersFalseParam(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "False Param",
		path:       "/users/:id",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
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
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	context.SetParamNames("id")
	context.SetParamValues("x")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllersTesting())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})
}

//test login success
func TestLoginGetUsersControllers(t *testing.T) {
	testCases := struct {
		name         string
		path         string
		expectStatus int
	}{

		name:         "Login Success",
		path:         "/users/:id",
		expectStatus: http.StatusOK,
	}

	e := InitEcho()
	InsertUser()
	body, error := json.Marshal(Login{Email: "andri@gmail.com", Password: "bismillah"})
	if error != nil {
		t.Error(t, error, "error")
	}
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/login")

	if assert.NoError(t, LoginUserControllers(context)) {

		var Users UserResponse
		res_body := res.Body.String()
		err := json.Unmarshal([]byte(res_body), &Users)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectStatus, res.Code)
		assert.Equal(t, testCases.name, Users.Message)

	}
}

//tes login error
func TestLoginUserControllersError(t *testing.T) {
	testCases := struct {
		name         string
		path         string
		expectStatus int
	}{

		name:         "Login Failed",
		path:         "/users/:id",
		expectStatus: http.StatusBadRequest,
	}

	e := InitEcho()
	config.DB.Migrator().DropTable(models.Users{})
	req := httptest.NewRequest(http.MethodPost, "/users/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	if assert.NoError(t, LoginUserControllers(context)) {

		var user UserResponse
		res_body := res.Body.String()
		err := json.Unmarshal([]byte(res_body), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectStatus, res.Code)
		assert.Equal(t, testCases.name, user.Message)

	}
}
