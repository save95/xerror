package ecode

type Repository interface {
	// FindByCode 通过 error code 搜索所有错误码定义
	FindByCode(code int) ([]Entity, error)
}
