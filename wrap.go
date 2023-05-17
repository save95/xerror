package xerror

import (
	"github.com/pkg/errors"
	"github.com/save95/xerror/xcode"
)

// Wrap 使用错误消息包装错误
func Wrap(err error, message string) *xError {
	// 如果没有错误，直接返回
	if nil == err {
		return nil
	}

	code := xcode.NewMessage(message)

	// 如果本身是 XError 则包装 error
	// 偷个懒，直接用 pkg/errors 来包装
	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
		code = xcode.NewWithMessage(xe.ErrorCode(), message)
	}

	return &xError{
		code:  code,
		error: errors.Wrap(err, message),
	}
}

// WrapWithCode 使用指定错误码包装错误
func WrapWithCode(err error, code int) *xError {
	return WrapWithXCode(err, xcode.New(code))
}

// WrapWithXCode 使用 xcode 来包装错误。
// 该方法响应的 httpStatus 和消息内容均来自 XCode，其 error 只在错误细节中展示。
// 在需要屏蔽错误细节，统一响应错误消息时使用。
func WrapWithXCode(err error, code xcode.XCode) *xError {
	// 如果没有错误，直接返回
	if nil == err {
		return nil
	}

	if nil == code {
		code = xcode.InternalServerError
	}

	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
	}

	return &xError{
		code:  code,
		error: errors.Wrap(err, code.String()),
	}
}

// WrapWithXCodeStatus 使用 XCode 的响应状态码来包装错误
// 该方法响应的 httpStatus 来自 XCode，但是响应消息内容来自 error
func WrapWithXCodeStatus(err error, code xcode.XCode) *xError {
	// 如果没有错误，直接返回
	if nil == err {
		return nil
	}

	if nil == code {
		code = xcode.InternalServerError
	}

	code = xcode.WithMessage(code, err.Error())

	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
	}

	return &xError{
		code:  code,
		error: err,
	}
}
