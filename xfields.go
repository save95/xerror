package xerror

type XFields interface {
	// 支持属性
	WithFields(fields ...interface{})
	// 获得属性
	GetFields() []interface{}
}
