package core

import (
	e2 "AynaAPI/api/core/e"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateApiResponseByStatus(status int, data map[string]interface{}) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: e2.GetMessage(status),
		Data:    data,
	}
}

func CreateEmptyApiResponseByStatus(status int) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: e2.GetMessage(status),
		Data:    nil,
	}
}
