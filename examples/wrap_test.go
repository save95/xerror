package examples

import (
	"errors"
	"fmt"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
)

func ExampleWrap() {
	err := errors.New("first error")
	werr := xerror.Wrap(err, "second error")
	fmt.Println(werr)

	// Output:
	// second error: first error
}

func ExampleWrap_xerror() {
	err := xerror.New("first xerror")
	werr := xerror.Wrap(err, "second xerror")

	fmt.Println(werr)
	fmt.Println(werr.ErrorCode() == err.ErrorCode())
	fmt.Println(werr.HttpStatus() == err.HttpStatus())

	// Output:
	// second xerror: first xerror
	// true
	// true
}

func ExampleWrapWithXCode() {
	err := xerror.New("gorm failed")
	werr := xerror.WrapWithXCode(err, xcode.DBFailed)

	fmt.Println(werr)
	fmt.Println(werr.ErrorCode() != err.ErrorCode())
	fmt.Println(werr.HttpStatus())

	// Output:
	// 数据库操作失败: gorm failed
	// true
	// 500
}

func ExampleWrapWithCode() {
	code := 3333
	err := xerror.New("first xerror")
	werr := xerror.WrapWithCode(err, code)

	fmt.Println(err.ErrorCode() != code)
	fmt.Println(werr)
	fmt.Println(werr.ErrorCode() == code)
	fmt.Println(werr.ErrorCode() != err.ErrorCode())
	fmt.Println(werr.HttpStatus())

	// Output:
	// true
	// 内部服务错误: first xerror
	// true
	// true
	// 500
}

func ExampleWrapWithCode_reserveCode() {
	code := 401
	err := xerror.New("first xerror")
	werr := xerror.WrapWithCode(err, code)

	fmt.Println(err.ErrorCode() != code)
	fmt.Println(werr)
	fmt.Println(werr.ErrorCode() == code)
	fmt.Println(werr.ErrorCode() == xcode.Unauthorized.Code())
	fmt.Println(werr.ErrorCode() != err.ErrorCode())
	fmt.Println(werr.HttpStatus())

	// Output:
	// true
	// 内部服务错误: first xerror
	// true
	// true
	// true
	// 401
}
