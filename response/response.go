package response

import (
	"net/http"
)

func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}

func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
	}
	return result
}

func AccessForbiddenResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Access Forbidden",
	}
	return result
}

func SuccessResponseData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    data,
	}
	return result
}

func SuccessResponseNonData() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	}
	return result
}

func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Login Failed",
	}
	return result
}

func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    data,
	}
	return result
}
