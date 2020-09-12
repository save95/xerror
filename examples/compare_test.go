package examples

import (
	"fmt"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
)

func ExampleIsXCode() {
	err := xerror.New("server error")
	success := xcode.InternalServerError
	fmt.Println(xerror.IsXCode(err, success))

	failed := xcode.BadGateway
	fmt.Println(xerror.IsXCode(err, failed))

	// Output: true
	// false
}

func ExampleIsErrorCode() {
	err := xerror.New("server error")
	fmt.Println(xerror.IsErrorCode(err, xcode.InternalServerError.Code()))

	// Output: true
}

func ExampleIsErrorCode_code() {
	code := 3333
	failed := 333
	err := xerror.WithCode(code, "server error")
	fmt.Println(xerror.IsErrorCode(err, code))
	fmt.Println(xerror.IsErrorCode(err, failed))

	// Output: true
	// false
}
