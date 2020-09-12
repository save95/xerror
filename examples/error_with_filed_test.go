package examples

import (
	"errors"
	"fmt"

	"github.com/save95/xerror"
)

func ExampleError_WithFields() {
	fmt.Printf("%#v\n", fieldError())

	// Output:
	// &xerror.Error{code:(*xcode.xCode)(0xc0000b23a0), error:(*errors.fundamental)(0xc0000b23e0), fields:[]interface {}{examples.field{Key1:"value1", Key2:-1, Key3:[]uint{0x0, 0x1, 0x2}}}}
}

func ExampleError_GetFields() {
	showFields(fieldError())
	showFields(errors.New("not fields"))

	// Output:
	// field: {Key1:value1 Key2:-1 Key3:[0 1 2]}
	// is not impl xerror.XFields
}

type field struct {
	Key1 string
	Key2 int
	Key3 []uint
}

func fieldError() error {
	err := xerror.New("field error")

	// 因为默认 xerror.XError 接口中并没有组合 xerror.XFields ，
	// 而，默认 xerror.Error 实现了 xerror.XFields 接口，
	// 所以，这里需要断言来使用
	if xew, ok := err.(xerror.XFields); ok {
		xew.WithFields(field{
			Key1: "value1",
			Key2: -1,
			Key3: []uint{0, 1, 2},
		})
	}

	return err
}

func showFields(err error) {
	// 获取 xerror.XFields 时，仍然需要使用断言
	if xe, ok := err.(xerror.XFields); ok {
		fields := xe.GetFields()
		for i := range fields {
			fmt.Printf("field: %+v\n", fields[i])
		}
	} else {
		fmt.Println("is not impl xerror.XFields")
	}
}
