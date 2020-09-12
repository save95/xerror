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
			ecodeEntity{ID: 1, Code: 500, Client: int(ecode.ClientDefault), Message: "default message: 500"},
			ecodeEntity{ID: 2, Code: 500, Client: int(ecode.ClientWeb), Message: "client web message: 500"},
			ecodeEntity{ID: 3, Code: 500, Client: int(ecode.ClientH5), Message: "client h5 message: 500"},
		}, nil
	case xcode.Forbidden.Code():
		return []ecode.Entity{
			ecodeEntity{ID: 4, Code: 403, Client: int(ecode.ClientApp), Message: "client app message: 403"},
		}, nil
	default:
		return []ecode.Entity{}, nil
	}
}

func ExampleError_ToMessage() {
	// 错误码仓库
	repo := &errorCodeRepository{}

	err := xerror.WithXCode(xcode.InternalServerError)
	// 已定义端口错误码消息，则按定义展示
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientWeb, Repository: repo}))
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientH5, Repository: repo}))
	// 没有定义端口错误码消息，则按默认端口消息展示
	fmt.Println(err.ToMessage(&ecode.Config{Client: ecode.ClientApp, Repository: repo}))

	err2 := xerror.WithXCode(xcode.Forbidden)
	// 已定义端口错误码消息，则按定义展示
	fmt.Println(err2.ToMessage(&ecode.Config{Client: ecode.ClientApp, Repository: repo}))
	// 没有定义端口错误码消息，也没有定义默认端口消息，则展示 XCode 消息
	fmt.Println(err2.ToMessage(&ecode.Config{Client: ecode.ClientH5, Repository: repo}))

	// Output:
	// client web message: 500
	// client h5 message: 500
	// default message: 500
	// client app message: 403
	// 无权访问
}
