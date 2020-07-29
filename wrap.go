package xerror

import (
	"github.com/pkg/errors"
	"github.com/save95/xerror/xcode"
)

func Wrap(err error, message string) XError {
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

	return &Error{
		code:  code,
		error: errors.Wrap(err, message),
	}
}

func WrapWithCode(err error, code int) XError {
	return WrapWithXCode(err, xcode.New(code))
}

func WrapWithXCode(err error, code xcode.XCode) XError {
	if nil == code {
		code = xcode.InternalServerError
	}

	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
	}

	// 如果 err 为 nil，只返回 xcode 错误
	if nil == err {
		err = errors.New(code.String())
	} else {
		err = errors.Wrap(err, code.String())
	}

	return &Error{
		code:  code,
		error: err,
	}
}
