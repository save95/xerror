package examples

import (
	"fmt"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
)

func ExampleNew() {
	err := xerror.New("whoops")
	fmt.Println(err)
	fmt.Println(err.HttpStatus())
	fmt.Println(err.ErrorCode())

	// Output: whoops
	// 500
	// 500
}

func ExampleWithCode() {
	code := 1001
	err := xerror.WithCode(code, "server error message")
	fmt.Println(err)
	fmt.Println(err.HttpStatus())
	fmt.Println(err.ErrorCode())

	// Output:
	// server error message
	// 500
	// 1001
}

func ExampleWithCode_reserveCode() {
	code := 500
	err := xerror.WithCode(code, "customize server error message")
	// 使用系统保留错误码，只替换错误码消息
	fmt.Println(err)
	// HTTP 状态码和错误码不会被覆写
	fmt.Println(err.HttpStatus())
	fmt.Println(err.ErrorCode())
	// 错误消息被覆写
	fmt.Println(err.String() != xcode.InternalServerError.String())

	// Output:
	// customize server error message
	// 500
	// 500
	// true
}

func ExampleWithCodef() {
	code := 1001
	err := xerror.WithCodef(code, "format error: code[%d] message: %s", code, "some error")
	fmt.Println(err)

	// Output:
	// format error: code[1001] message: some error
}

func ExampleWithXCode() {
	err := xerror.WithXCode(xcode.InternalServerError)
	fmt.Println(err)

	// Output:
	// 内部服务错误
}

func ExampleWithXCodeMessage() {
	err := xerror.WithXCodeMessage(xcode.InternalServerError, "customize error message")
	fmt.Println(err)

	// Output:
	// customize error message
}

func ExampleWithXCodeMessagef() {
	err := xerror.WithXCodeMessagef(xcode.InternalServerError, "format error message: %s", "some error")
	fmt.Println(err)

	// Output:
	// format error message: some error
}
