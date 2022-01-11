package xerror

import "fmt"

// github.com/pkg/errors 包兼容函数

func Errorf(format string, args ...interface{}) *xError {
	return New(fmt.Sprintf(format, args...))
}

func Wrapf(err error, format string, args ...interface{}) *xError {
	if nil == err {
		return nil
	}

	return Wrap(err, fmt.Sprintf(format, args...))
}

func WithMessage(err error, message string) *xError {
	if nil == err {
		return nil
	}

	return Wrap(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) *xError {
	if nil == err {
		return nil
	}

	return Wrap(err, fmt.Sprintf(format, args...))
}
