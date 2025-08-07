package handler

import internalError "github.com/configservice/internal/error"

type ResponseInterface interface {
	SetOk(data interface{}) HTTPResponse
	SetOkWithStatus(status int, data interface{}) HTTPResponse
	ImportJSONWrapError(err *internalError.JSONWrapError) HTTPResponse
	HasError() bool
	GetData() interface{}
	GetStatus() int
	GetErrCode() int
	GetErrorMessage() string
	HasNoContent() bool
}
