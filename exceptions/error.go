package exceptions

import (
	"gin-api/utils"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	return &Error{
		code,
		message,
	}
}

func Error500(err error) {
	if err != nil {
		utils.SpewDump(err.Error())
		panic(NewError(http.StatusServiceUnavailable, err.Error()))
	}
}

func Error422(err error) {
	if err != nil {
		panic(NewError(http.StatusUnprocessableEntity, err.Error()))
	}
}

func Error404(err error) {
	if err != nil {
		panic(NewError(http.StatusNotFound, "记录不存在"))
	}
}

func Error401(err error) {
	if err != nil {
		panic(NewError(http.StatusUnauthorized, err.Error()))
	}
}
