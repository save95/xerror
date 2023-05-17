package xerror

import "github.com/save95/xerror/xcode"

// IsXCode 判断错误是否为 XCode。只要其错误码一致，不管错误消息是否一致，均则视为一致
func IsXCode(err error, code xcode.XCode) bool {
	return IsErrorCode(err, code.Code())
}

// IsErrorCode 判断错误的错误码是否一致
func IsErrorCode(err error, code int) bool {
	if xe, ok := err.(XError); ok {
		return xe.ErrorCode() == code
	}

	return false
}
