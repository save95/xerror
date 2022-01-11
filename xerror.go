package xerror

import "github.com/save95/xerror/ecode"

type XError interface {
	// String 获得错误的消息
	// 该消息为其包裹的 `XCode` 的 `String()` 的返回值；
	// 如果没有包裹 `XCode`，则返回固定字符串 `error`
	String() string
	// Error 获得 error 栈消息
	Error() string
	// Unwrap 解包获得 error
	Unwrap() error
	// HttpStatus 获得错误的 HTTP STATUS
	HttpStatus() int
	// ErrorCode 获得错误码
	ErrorCode() int
	// ToMessage 将错误转成消息
	ToMessage(config *ecode.Config) string
}
