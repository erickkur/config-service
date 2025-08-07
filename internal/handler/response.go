package handler

import (
	"fmt"
	"net/http"

	internalError "github.com/configservice/internal/error"
	"github.com/configservice/internal/logger"
)

type HTTPResponse struct {
	status    int
	data      interface{}
	err       error
	errCode   int
	message   string
	noContent bool
	Logger    logger.Interface
}

type Dep struct {
	Log logger.Interface
}

func NewResponse(dep Dep) *HTTPResponse {
	return &HTTPResponse{
		Logger: dep.Log,
	}
}

// SetOk ...
func (res HTTPResponse) SetOk(data interface{}) HTTPResponse {
	res.status = http.StatusOK
	res.data = data
	res.message = "success"

	return res
}

// SetOkWithStatus ...
func (res HTTPResponse) SetOkWithStatus(status int, data interface{}) HTTPResponse {
	res.status = status
	res.data = data
	res.message = "success"
	res.noContent = status == http.StatusNoContent

	return res
}

// ImportJSONWrapError ...
func (res HTTPResponse) ImportJSONWrapError(err *internalError.JSONWrapError) HTTPResponse {
	res.status = err.Status
	res.err = err.Error
	res.errCode = err.Code
	res.message = err.Message

	errStr := fmt.Sprintf(
		"Error occured. %s",
		err.StringWithError(),
	)

	if err.Status >= 500 {
		res.Logger.Error(errStr)
	}

	return res
}

// HasError ...
func (res HTTPResponse) HasError() bool {
	return res.err != nil
}

// GetData ...
func (res HTTPResponse) GetData() interface{} {
	return res.data
}

// GetStatus ...
func (res HTTPResponse) GetStatus() int {
	if res.status != 0 {
		return res.status
	}
	return http.StatusInternalServerError
}

// GetErrCode ...
func (res HTTPResponse) GetErrCode() int {
	if res.errCode != 0 {
		return res.errCode
	}

	return 0
}

// GetErrorMessage get error message from message or error object
func (res HTTPResponse) GetErrorMessage() string {
	if res.message != "" {
		return res.message
	}

	return res.err.Error()
}

// HasNoContent ...
func (res HTTPResponse) HasNoContent() bool {
	return res.noContent
}
