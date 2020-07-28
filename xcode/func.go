package xcode

import (
	"net/http"
)

func New(code int) XCode {
	// 如果是已定义的 code，直接返回；否则构建一个
	for _, c := range allCode {
		if c.code == code {
			return c
		}
	}

	return NewWithMessage(code, InternalServerError.message)
}

func NewMessage(message string) XCode {
	// copy InternalServerError
	c := &xCode{}
	*c = *InternalServerError

	c.message = message

	return c
}

func NewWithMessage(code int, message string) XCode {
	// 检查错误码
	if MinReservedCode <= code && code <= MaxReservedCode {
		return ErrorCodeFailed
	}

	return &xCode{
		httpStatus: http.StatusInternalServerError,
		code:       code,
		message:    message,
	}
}

func WithMessage(xcode XCode, message string) XCode {
	return &xCode{
		httpStatus: xcode.HttpStatus(),
		code:       xcode.Code(),
		message:    message,
	}
}
