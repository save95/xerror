package xerror

import (
	"encoding/json"
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

func ParsePayload(err error) string {
	if xf, ok := err.(XFields); ok {
		fields := xf.GetFields()
		if fields != nil && len(fields) > 0 {
			xfbs, _ := json.Marshal(fields)
			return string(xfbs)
		}
	}

	return ""
}

// FormatStackTrace 格式化错误栈
func FormatStackTrace(err error) string {
	return fmt.Sprintf("%+v", err)
}
