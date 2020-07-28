package ecode

// 错误码实体
// 唯一建：Code + Client
type Entity interface {
	GetCode() int       // 错误码
	GetClient() int     // 错误显示的端口（客户端）
	GetHttpStatus() int // HTTP 状态码
	GetMessage() string // 错误消息
}
