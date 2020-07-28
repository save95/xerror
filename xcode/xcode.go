package xcode

type XCode interface {
	// 错误码
	Code() int
	// HTTP 状态码，参考：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status
	HttpStatus() int
	// 错误码消息
	String() string
}
