package app

import "AynaAPI/server/app/e"

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Detail  []string `json:"detail"`
}

func (e *ErrorResponse) WithDetails(details ...string) *ErrorResponse {
	newError := *e
	newError.Detail = []string{}
	for _, d := range details {
		newError.Detail = append(newError.Detail, d)
	}
	return &newError
}

func (g *AppGin) MakeErrorResponse(httpCode int, statusCode int, details ...string) {
	g.C.JSON(httpCode, ErrorResponse{
		Code:    statusCode,
		Message: e.GetMessage(statusCode),
		Detail:  details,
	})
	return
}
