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
		Unauthorized = xCode{httpStatus: http.StatusUnauthorized, code: 401, message: "overwrite 请求未授权"}
		Forbidden    = xCode{httpStatus: http.StatusForbidden, code: 403, message: "overwrite 无权访问"}
	)

	var allCode = []xcode.XCode{
		Unauthorized,
		Forbidden,
	}

	xerror.LoadCodes(allCode...)

	fmt.Println(xcode.Unauthorized.String())
	fmt.Println(xerror.WithXCode(Unauthorized))
	fmt.Println(xerror.WithXCode(Forbidden))

	// Output:
	// 请求未授权
	// overwrite 请求未授权
	// overwrite 无权访问
}

func ExampleAppendCodes() {
	var (
		E1001 = xCode{httpStatus: http.StatusInternalServerError, code: 1001, message: "1001 error"}
		E1002 = xCode{httpStatus: http.StatusInternalServerError, code: 1002, message: "1002 error"}
	)

	xerror.AppendCodes(E1001, E1002)

	fmt.Println(xerror.WithXCode(E1001))
	fmt.Println(xerror.WithXCode(E1002))

	err := xerror.WithCode(E1001.code, "with message")
	fmt.Println(err)
	fmt.Println(err.ErrorCode() == E1001.code)
	fmt.Println(err.String())
	fmt.Println(err.String() != E1001.message)

	// Output:
	// 1001 error
	// 1002 error
	// with message
	// true
	// with message
	// true
}
