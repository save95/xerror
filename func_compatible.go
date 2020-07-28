// github.com/pkg/errors 包兼容函数
package xerror

import "fmt"

func Errorf(format string, args ...interface{}) error {
	return New(fmt.Sprintf(format, args...))
}

func Wrapf(err error, format string, args ...interface{}) error {
	if nil == err {
		return nil
	}

	return Wrap(err, fmt.Sprintf(format, args...))
}

func WithMessage(err error, message string) error {
	if nil == err {
		return nil
	}

	return Wrap(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	if nil == err {
		return nil
	}

	return Wrap(err, fmt.Sprintf(format, args...))
}
