package xcode

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRepository_LoadCodes(t *testing.T) {
	var (
		InternalServerError = &xCode{httpStatus: http.StatusInternalServerError, code: 500, message: "内部服务错误1"}
	)

	rep := Repository()
	fmt.Printf("%+v\n", _repo)
	rep.LoadCodes(InternalServerError) // 重新加载错误码

	xc := New(500)
	fmt.Printf("%+v\n", _repo)
	fmt.Printf("%#v\n", xc)
}

func TestRepository_AppendCodes(t *testing.T) {
	var (
		InternalServerError  = &xCode{httpStatus: http.StatusInternalServerError, code: 500, message: "内部服务错误2"}
		InternalServerError2 = &xCode{httpStatus: http.StatusInternalServerError, code: 5000, message: "内部服务错误3"}
	)

	rep := Repository()
	fmt.Printf("%+v\n", _repo)
	rep.AppendCodes(InternalServerError, InternalServerError2)

	fmt.Printf("%#v\n", New(500))  // 不会覆写
	fmt.Printf("%#v\n", New(5000)) // 追加
}
