package xcode

import "net/http"

const (
	MinReservedCode = -1000 // 系统保留错误码最低值
	MaxReservedCode = 1000  // 系统保留错误码最大值
	ForbiddenCode   = 0     // 禁止的错误码
)

// 系统保留错误码定义
var (
	ErrorCodeFailed = &xCode{httpStatus: http.StatusInternalServerError, code: 100, message: "错误码定义错误"}

	RequestParamError   = &xCode{httpStatus: http.StatusBadRequest, code: 400, message: "请求参数错误"}
	Unauthorized        = &xCode{httpStatus: http.StatusUnauthorized, code: 401, message: "请求未授权"}
	Forbidden           = &xCode{httpStatus: http.StatusForbidden, code: 403, message: "无权访问"}
	ResourceNotFound    = &xCode{httpStatus: http.StatusNotFound, code: 404, message: "请求资源不存在"}
	MethodNotAllowed    = &xCode{httpStatus: http.StatusMethodNotAllowed, code: 405, message: "请求方法不支持"}
	NotAcceptable       = &xCode{httpStatus: http.StatusNotAcceptable, code: 406, message: "不接受的请求"}
	RequestTimeout      = &xCode{httpStatus: http.StatusRequestTimeout, code: 408, message: "请求超时"}
	InternalServerError = &xCode{httpStatus: http.StatusInternalServerError, code: 500, message: "内部服务错误"}
	NotImplemented      = &xCode{httpStatus: http.StatusNotImplemented, code: 501, message: "请求方法未实现"}
	BadGateway          = &xCode{httpStatus: http.StatusBadGateway, code: 502, message: "网关异常"}
	GatewayTimeout      = &xCode{httpStatus: http.StatusGatewayTimeout, code: 504, message: "网关超时"}
	VersionNotSupported = &xCode{httpStatus: http.StatusHTTPVersionNotSupported, code: 505, message: "不支持的版本"}

	// DB 错误，前缀 1

	DBFailed            = &xCode{httpStatus: http.StatusInternalServerError, code: 1001, message: "数据库操作失败"}
	DBTransactionError  = &xCode{httpStatus: http.StatusInternalServerError, code: 1002, message: "数据库事务错误"}
	DBRecordNotFound    = &xCode{httpStatus: http.StatusNotFound, code: 1003, message: "无相关记录"}
	DBRecordExist       = &xCode{httpStatus: http.StatusInternalServerError, code: 1004, message: "数据已存在"}
	DBRequestParamError = &xCode{httpStatus: http.StatusBadRequest, code: 1005, message: "数据库操作请求参数错误"}
	DBDuplicateEntry    = &xCode{httpStatus: http.StatusBadRequest, code: 1006, message: "数据库存在重复数据"}

	// GRPC 错误，前缀 2

	GRPCFailed         = &xCode{httpStatus: http.StatusInternalServerError, code: 2001, message: "GRPC 服务错误"}
	GRPCMethodNotFound = &xCode{httpStatus: http.StatusInternalServerError, code: 2002, message: "不支持的 GRPC 方法"}
	GRPCUnauthorized   = &xCode{httpStatus: http.StatusInternalServerError, code: 2003, message: "GRPC 访问未授权"}
	GRPCAccessDenied   = &xCode{httpStatus: http.StatusInternalServerError, code: 2004, message: "GRPC 访问权限不足"}
	GRPCCanceled       = &xCode{httpStatus: http.StatusInternalServerError, code: 2005, message: "GRPC 客户端取消请求"}
	GRPCTimeout        = &xCode{httpStatus: http.StatusInternalServerError, code: 2006, message: "GRPC 处理超时"}
)

var allCode = []XCode{
	ErrorCodeFailed,

	Unauthorized,
	Forbidden,
	ResourceNotFound,
	MethodNotAllowed,
	NotAcceptable,
	RequestParamError,
	InternalServerError,
	NotImplemented,
	RequestTimeout,
	BadGateway,
	GatewayTimeout,
	VersionNotSupported,

	DBFailed,
	DBTransactionError,
	DBRecordNotFound,
	DBRecordExist,
	DBRequestParamError,
	DBDuplicateEntry,

	GRPCFailed,
	GRPCMethodNotFound,
	GRPCUnauthorized,
	GRPCAccessDenied,
	GRPCCanceled,
	GRPCTimeout,
}
