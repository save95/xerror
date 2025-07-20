package xerror

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/save95/xerror/ecode"
	"github.com/save95/xerror/xcode"
)

type xError struct {
	code   xcode.XCode
	error  error
	fields []interface{}
}

func (e *xError) String() string {
	if e.code == nil {
		return xcode.InternalServerError.String()
	}

	return e.code.String()
}

func (e *xError) Error() string {
	if e.error == nil {
		return xcode.InternalServerError.String()
	}

	return e.error.Error()
}

func (e *xError) Unwrap() error {
	return e.error
}

func (e *xError) HttpStatus() int {
	if e.code == nil {
		return xcode.InternalServerError.HttpStatus()
	}

	return e.code.HttpStatus()
}

func (e *xError) ErrorCode() int {
	if e.code == nil {
		return xcode.InternalServerError.Code()
	}

	return e.code.Code()
}

func (e *xError) ToMessage(config *ecode.Config) string {
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

func (e *xError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			// 展示 xField 内容
			if e.fields != nil && len(e.fields) > 0 {
				if bs, err := json.Marshal(e.fields); nil == err {
					_, _ = fmt.Fprintf(s, "\nfields: %s\n", string(bs))
				}
			}
			_, _ = fmt.Fprintf(s, "%+v", e.error)
			return
		}
		fallthrough
	case 's':
		if s.Flag('-') {
			_, _ = io.WriteString(s, e.String())
			return
		}
		_, _ = fmt.Fprintf(s, "[%d] %s", e.ErrorCode(), e.error)
	}
}

func (e *xError) WithFields(field ...interface{}) *xError {
	if e.fields == nil {
		e.fields = make([]interface{}, 0)
	}
	e.fields = append(e.fields, field...)

	return e
}

func (e *xError) GetFields() []interface{} {
	if e.fields == nil {
		return make([]interface{}, 0)
	}
	return e.fields
}
