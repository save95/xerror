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

func TestWrap_Field(t *testing.T) {
	type field struct {
		Key1 string
		Key2 int
		Key3 []uint
	}

	err := New("field error").
		WithFields(field{
			Key1: "value1",
			Key2: -1,
			Key3: []uint{0, 1, 2},
		}).
		WithFields("append field").
		WithFields(map[string]interface{}{
			"mapKey1": "value1",
			"mapKey2": 2,
			"mapKey3": []interface{}{0, "a", []int8{1, 2}},
		})

	wxe := Wrap(err, "wrap test")

	fmt.Printf("%+v\n", wxe)
	fmt.Printf("%+v\n", wxe.Error())  // 含有错误的栈消息
	fmt.Printf("%+v\n", wxe.String()) // 返回错误的消息，即 unwrap 之后的消息
	fmt.Printf("fileds: %#v\n", wxe.GetFields())
}

func TestWrap_Trace(t *testing.T) {
	var a1 = func() error {
		return Errorf("origin error")
	}

	var a2 = func() error {
		if err := a1(); nil != err {
			return Wrap(err, "wrap a2 error")
		}
		return nil
	}

	var a3 = func() error {
		if err := a2(); nil != err {
			return Wrap(err, "wrap a3 error")
		}
		return nil
	}

	res := a3()

	fmt.Printf("%-s\n", res)
	fmt.Printf("%s\n", res)
	fmt.Printf("%+v\n", res)
}
