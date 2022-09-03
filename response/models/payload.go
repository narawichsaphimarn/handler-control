package models

type ResponseSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

type ResponseUnsuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail"`
	Body    any    `json:"body"`
}
