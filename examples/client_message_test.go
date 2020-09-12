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
	switch code {
	case xcode.InternalServerError.Code():
		return []ecode.Entity{
			ecodeEntity{ID: 1, Code: 500, Client: int(ecode.ClientWeb), Message: "client web message: 500"},
			ecodeEntity{ID: 2, Code: 500, Client: int(ecode.ClientH5), Message: "client h5 message: 500"},
			ecodeEntity{ID: 3, Code: 500, Client: int(ecode.ClientApp), Message: "client app message: 500"},
			ecodeEntity{ID: 4, Code: 500, Client: int(ecode.ClientWechatPublicAccount), Message: "client wechat message: 500"},
		}, nil
	case xcode.Forbidden.Code():
		return []ecode.Entity{
			ecodeEntity{ID: 5, Code: 403, Client: int(ecode.ClientDefault), Message: "default message: 403"},
			ecodeEntity{ID: 6, Code: 403, Client: int(ecode.ClientApp), Message: "client app message: 403"},
		}, nil
	default:
		return []ecode.Entity{}, nil
	}
}

func ExampleError_ToMessage() {
	repo := &errorCodeRepository{}

	err := xerror.WithXCode(xcode.InternalServerError)
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientWechatPublicAccount, Repository: repo}))
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientWechatMiniProgram, Repository: repo}))

	err2 := xerror.WithXCode(xcode.Forbidden)
	fmt.Println(err2.ToMessage(&ecode.Config{Client: ecode.ClientApp, Repository: repo}))
	fmt.Println(err2.ToMessage(&ecode.Config{Client: ecode.ClientH5, Repository: repo}))

	// Output:
	// client wechat message: 500
	// 内部服务错误
	// client app message: 403
	// default message: 403
}
