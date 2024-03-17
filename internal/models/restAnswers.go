package models

type SuccessResponse struct {
	Message interface{} `json:"msg"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Error interface{} `json:"err"`
}
