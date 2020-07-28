package xerror

import (
	"github.com/save95/xerror/ecode"
	"github.com/save95/xerror/xcode"
)

type Error struct {
	code   xcode.XCode
	error  error
	fields []interface{}
}

func (e *Error) String() string {
	if e.code == nil {
		return "error"
	}

	return e.code.String()
}

func (e *Error) Error() string {
	if e.error == nil {
		return "error"
	}

	return e.error.Error()
}

func (e *Error) Unwrap() error {
	return e.error
}

func (e *Error) HttpStatus() int {
	if e.code == nil {
		return xcode.InternalServerError.Code()
	}

	return e.code.HttpStatus()
}

func (e *Error) ErrorCode() int {
	return e.code.Code()
}

func (e *Error) ToMessage(config *ecode.Config) string {
	msg := e.String()
	if nil == config || nil == config.Repository {
		return msg
	}

	// 从资源仓库获取所有端口的解析错误码解析规则
	if es, err := config.Repository.FindByCode(e.code.Code()); nil == err {
		defaultMsg := msg
		for _, entity := range es {
			if entity.GetCode() != e.code.Code() {
				break
			}

			// 取得默认消息
			if entity.GetClient() == int(ecode.ClientDefault) {
				defaultMsg = entity.GetMessage()
			}

			if entity.GetClient() == int(config.Client) && len(entity.GetMessage()) > 0 {
				msg = entity.GetMessage()
				break
			}
		}

		// 如果没有对应端口消息，但是有默认消息，则使用默认消息
		if len(msg) == 0 && len(defaultMsg) > 0 {
			msg = defaultMsg
		}
	}

	return msg
}

func (e *Error) WithFields(field ...interface{}) {
	if e.fields == nil {
		e.fields = make([]interface{}, 0)
	}
	e.fields = append(e.fields, field...)
}

func (e *Error) GetFields() []interface{} {
	if e.fields == nil {
		return make([]interface{}, 0)
	}
	return e.fields
}

func (e *Error) LoadCodes(codes ...xcode.XCode) {
	xcode.Repository().LoadCodes(codes...)
}

func (e *Error) AppendCodes(codes ...xcode.XCode) {
	xcode.Repository().AppendCodes(codes...)
}
