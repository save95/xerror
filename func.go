package xerror

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/save95/xerror/xcode"
)

func New(message string) *xError {
	return &xError{
		code:  xcode.NewMessage(message),
		error: errors.New(message),
	}
}

func WithCode(code int, message string) *xError {
	return &xError{
		code:  xcode.NewWithMessage(code, message),
		error: errors.New(message),
	}
}

func WithCodef(code int, format string, args ...interface{}) *xError {
	return WithCode(code, fmt.Sprintf(format, args...))
}

func WithXCode(code xcode.XCode) *xError {
	return &xError{
		code:  code,
		error: errors.New(code.String()),
	}
}

func WithXCodeMessage(code xcode.XCode, message string) *xError {
	if len(message) > 0 {
		code = xcode.WithMessage(code, message)
	}

	return &xError{
		code:  code,
		error: errors.New(code.String()),
	}
}

func WithXCodeMessagef(code xcode.XCode, format string, args ...interface{}) *xError {
	return WithXCodeMessage(code, fmt.Sprintf(format, args...))
}

func LoadCodes(codes ...xcode.XCode) {
	xcode.Repository().LoadCodes(codes...)
}

func AppendCodes(codes ...xcode.XCode) {
	xcode.Repository().AppendCodes(codes...)
}

func IsXCode(err error, code xcode.XCode) bool {
	return IsErrorCode(err, code.Code())
}

func IsErrorCode(err error, code int) bool {
	if xe, ok := err.(XError); ok {
		return xe.ErrorCode() == code
	}

	return false
}
