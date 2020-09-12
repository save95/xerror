package examples

import (
	"fmt"

	"github.com/save95/xerror"
	"github.com/save95/xerror/ecode"
	"github.com/save95/xerror/xcode"
)

type ecodeEntity struct {
	ID         uint
	Code       int
	Client     int
	HttpStatus int
	Message    string
}

func (e ecodeEntity) GetCode() int {
	return e.Code
}

func (e ecodeEntity) GetClient() int {
	return e.Client
}

func (e ecodeEntity) GetHttpStatus() int {
	return e.HttpStatus
}

func (e ecodeEntity) GetMessage() string {
	return e.Message
}

// 消息
type errorCodeRepository struct{}

func (er errorCodeRepository) FindByCode(code int) ([]ecode.Entity, error) {
	if code != xcode.InternalServerError.Code() {
		return []ecode.Entity{}, nil
	}

	return []ecode.Entity{
		ecodeEntity{ID: 1, Code: code, Client: int(ecode.ClientWeb), Message: "client web message"},
		ecodeEntity{ID: 2, Code: code, Client: int(ecode.ClientH5), Message: "client h5 message"},
		ecodeEntity{ID: 3, Code: code, Client: int(ecode.ClientApp), Message: "client app message"},
		ecodeEntity{ID: 4, Code: code, Client: int(ecode.ClientWechatPublicAccount), Message: "client wechat message"},
	}, nil
}

func ExampleError_ToMessage() {
	err := xerror.WithXCode(xcode.InternalServerError)
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientWechatPublicAccount, Repository: &errorCodeRepository{}}))
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientWechatMiniProgram, Repository: &errorCodeRepository{}}))

	// Output:
	// client wechat message
	// 内部服务错误
}
