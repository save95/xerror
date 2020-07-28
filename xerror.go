package xerror

import "github.com/save95/xerror/ecode"

type XError interface {
	// 获得 code 的消息
	String() string
	// 获得 error 栈消息
	Error() string
	// 解包获得 error
	Unwrap() error
	// 获得错误的 HTTP STATUS
	HttpStatus() int
	// 错误 code
	ErrorCode() int
	// 将错误转成消息
	ToMessage(config *ecode.Config) string
}
