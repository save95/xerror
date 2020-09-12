package examples

import (
	"fmt"
	"net/http"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
)

type xCode struct {
	httpStatus int
	code       int
	message    string
}

func (c xCode) Code() int {
	return c.code
}

func (c xCode) HttpStatus() int {
	return c.httpStatus
}

func (c xCode) String() string {
	return c.message
}

func ExampleLoadCodes() {
	var (
		unauthorized = xCode{httpStatus: http.StatusInternalServerError, code: 401, message: "overwrite 请求未授权"}
		forbidden    = xCode{httpStatus: http.StatusForbidden, code: 403, message: "overwrite 无权访问"}
	)

	var allCode = []xcode.XCode{
		unauthorized,
		forbidden,
	}

	// 加载自定义的错误码列表，会通过 xcode.XCode.Code() 值进行覆写
	xerror.LoadCodes(allCode...)

	// 系统错误码的消息
	fmt.Println(xcode.Unauthorized.String())
	fmt.Println(xcode.Forbidden.String())
	// 错误消息被覆写
	fmt.Println(xerror.WithCode(unauthorized.code, unauthorized.message))
	fmt.Println(xerror.WithCode(forbidden.code, forbidden.message))
	// HTTP 状态码被覆写
	fmt.Println(xerror.WithCode(unauthorized.code, "overwrite").HttpStatus())

	// Output:
	// 请求未授权
	// 无权访问
	// overwrite 请求未授权
	// overwrite 无权访问
	// 500
}

func ExampleAppendCodes() {
	var (
		E403  = xCode{httpStatus: http.StatusInternalServerError, code: 403, message: "overwrite 403 error"}
		E1001 = xCode{httpStatus: http.StatusUnauthorized, code: 1001, message: "1001 error"}
	)

	// 追加系统错误码列表；如果已存在，则直接跳过
	xerror.AppendCodes(E403, E1001)

	// 追加后，可以直接使用
	fmt.Println(xerror.WithXCode(E403))
	fmt.Println(xerror.WithXCode(E403).HttpStatus())
	fmt.Println(xerror.WithXCode(E1001))

	// 对于已存在的错误码，无法被覆写
	// 这时，调用 xerror.WithCode 所使用的为系统定义的错误码
	err := xerror.WithCode(E403.code, "with message")
	fmt.Println(err)
	fmt.Println(err.ErrorCode() == E403.code)
	fmt.Println(err.HttpStatus() != E403.httpStatus)
	fmt.Println(err.String())
	fmt.Println(err.String() != E403.message)

	// Output:
	// overwrite 403 error
	// 500
	// 1001 error
	// with message
	// true
	// true
	// with message
	// true
}
