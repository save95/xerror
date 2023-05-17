package ecode

// Entity 错误码实体
// 唯一建：Code + Client
type Entity interface {
	// GetCode 错误码
	GetCode() int
	// GetClient 错误显示的端口（客户端）
	GetClient() int
	// GetHttpStatus HTTP 状态码
	GetHttpStatus() int
	// GetMessage 错误消息
	GetMessage() string
}
