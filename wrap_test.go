package xerror

import (
	"fmt"
	"testing"

	"github.com/save95/xerror/xcode"
)

func TestWrap(t *testing.T) {
	xc := xcode.NewWithMessage(10001, "某种内容错误")
	xe := WithXCode(xc)
	wxe := Wrap(xe, "wrap test")

	fmt.Printf("%+v\n", wxe)
	fmt.Printf("%+v\n", wxe.Error())  // 含有错误的栈消息
	fmt.Printf("%+v\n", wxe.String()) // 返回错误的消息，即 unwrap 之后的消息
}
