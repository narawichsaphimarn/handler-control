package services

import "github.com/narawichsaphimarn/handlercontrol/response/models"

func MapResponseSuccess(code int, msg string, body any) *models.ResponseSuccess {
	payload := models.ResponseSuccess{}
	payload.Code = code
	payload.Message = msg
	payload.Body = body
	return &payload
}

func MapResponseUnsuccess(code int, msg string, detail string, body any) *models.ResponseUnsuccess {
	payload := models.ResponseUnsuccess{}
	payload.Code = code
	payload.Message = msg
	payload.Detail = detail
	payload.Body = body
	return &payload
}
