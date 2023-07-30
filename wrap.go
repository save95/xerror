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

	// 如果本身是 XError 则使用其错误码
	code := xcode.NewMessage(message)
	if xe, ok := err.(XError); ok {
		code = xcode.NewWithMessage(xe.ErrorCode(), message)
	}

	return &xError{
		code:   code,
		error:  errors.Wrap(getOriginError(err), message),
		fields: getFields(err),
	}
}

// 获取原始错误，防止过度包装
func getOriginError(err error) error {
	if xe, ok := err.(XError); ok {
		return xe.Unwrap()
	}
	return err
}

// 获取 xError 的自定义字段，方便追加到包装后的错误中
func getFields(err error) []interface{} {
	if xe, ok := err.(XFields); ok {
		return xe.GetFields()
	}
	return nil
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

	return &xError{
		code:   code,
		error:  errors.Wrap(getOriginError(err), code.String()),
		fields: getFields(err),
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

	return &xError{
		code:   code,
		error:  getOriginError(err),
		fields: getFields(err),
	}
}
