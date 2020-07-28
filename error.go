package xerror

import (
	"github.com/save95/xerror/ecode"
	"github.com/save95/xerror/xcode"
)

type Error struct {
	code   xcode.XCode
	error  error
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
	var defaultMsg, clientMsg string

	// 从资源仓库获取所有端口的解析错误码解析规则
	if config != nil && config.Repository != nil {
		if es, err := config.Repository.FindByCode(e.code.Code()); nil == err {
			for _, entity := range es {
				if entity.GetCode() != e.code.Code() {
					break
				}

				// 取得默认消息
				if entity.GetClient() == int(ecode.ClientDefault) {
					defaultMsg = entity.GetMessage()
				}

				if entity.GetClient() == int(config.Client) && len(entity.GetMessage()) > 0 {
					clientMsg = entity.GetMessage()
					break
				}
			}
		}
	}

	// 如果找到对应端口消息，则返回端口消息
	// 否则，查看是否有端口默认消息，如果有，则返回
	// 否则，查看是否有 XCode 消息，如果有，则返回
	// 否则，返回错误消息
	if len(clientMsg) > 0 {
		return clientMsg
	} else if len(defaultMsg) > 0 {
		return defaultMsg
	} else if len(e.code.String()) > 0 {
		return e.code.String()
	} else {
		return e.Error()
	}
}
