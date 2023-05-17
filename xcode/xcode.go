package xcode

type XCode interface {
	// Code 错误码
	Code() int
	// HttpStatus HTTP 状态码，参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status
	HttpStatus() int
	// String 错误码消息
	String() string
}
