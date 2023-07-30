package examples

import (
	"errors"
	"fmt"

	"github.com/save95/xerror"
)

func ExampleError_WithFields() {
	showFields(fieldError())
	showFields(errors.New("not fields"))

	// Output:
	// field: {Key1:value1 Key2:-1 Key3:[0 1 2]}
	// field: append
	// field: {Key1:value2 Key2:1 Key3:[2 1 2]}
	// is not impl xerror.XFields
}

func ExampleError_WithFields_Wrap() {
	showFields(xerror.Wrap(fieldError(), "wrap field error"))

	//Output:
	//field: {Key1:value1 Key2:-1 Key3:[0 1 2]}
	//field: append
	//field: {Key1:value2 Key2:1 Key3:[2 1 2]}
}

type field struct {
	Key1 string
	Key2 int
	Key3 []uint
}

func fieldError() error {
	return xerror.New("field error").
		WithFields(field{
			Key1: "value1",
			Key2: -1,
			Key3: []uint{0, 1, 2},
		}).
		WithFields("append").
		WithFields(field{
			Key1: "value2",
			Key2: 1,
			Key3: []uint{2, 1, 2},
		})
}

func showFields(err error) {
	// 获取 xerror.XFields 时，需要使用断言
	if xe, ok := err.(xerror.XFields); ok {
		fields := xe.GetFields()
		for i := range fields {
			fmt.Printf("field: %+v\n", fields[i])
		}
	} else {
		fmt.Println("is not impl xerror.XFields")
	}
}
