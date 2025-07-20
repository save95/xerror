package examples

import (
	"errors"
	"fmt"

	"github.com/save95/xerror"
)

func ExampleError_unwrap() {
	err := xerror.New("xerror")
	fmt.Printf("%+v\n", err.Unwrap())

	// Output:
	// xerror
	// github.com/save95/xerror.New
	// 	/home/user/dev/xerror/func.go:13
	// github.com/save95/xerror/examples.ExampleError_Unwrap
	// 	/home/user/dev/xerror/examples/show_stack_test.go:11
	// testing.runExample
	// 	/usr/local/go/src/testing/run_example.go:62
	// testing.runExamples
	// 	/usr/local/go/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/go/src/testing/testing.go:1118
	// main.main
	// 	_testmain.go:88
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:203
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1357
}

func ExampleError_unwrap_with_wrap() {
	err := xerror.New("xerror")
	werr := xerror.Wrap(err, "wrap message")
	fmt.Printf("%+v\n", werr.Unwrap())

	// Output:
	// xerror
	// github.com/save95/xerror.New
	// 	/home/user/dev/xerror/func.go:13
	// github.com/save95/xerror/examples.ExampleError_unwrap_with_wrap
	// 	/home/user/dev/xerror/examples/show_stack_test.go:35
	// testing.runExample
	// 	/usr/local/go/src/testing/run_example.go:62
	// testing.runExamples
	// 	/usr/local/go/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/go/src/testing/testing.go:1118
	// main.main
	// 	_testmain.go:88
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:203
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1357
	// wrap message
	// github.com/save95/xerror.Wrap
	// 	/home/user/dev/xerror/wrap.go:25
	// github.com/save95/xerror/examples.ExampleError_unwrap_with_wrap
	// 	/home/user/dev/xerror/examples/show_stack_test.go:36
	// testing.runExample
	// 	/usr/local/go/src/testing/run_example.go:62
	// testing.runExamples
	// 	/usr/local/go/src/testing/example.go:44
	// testing.(*M).Run
	// 	/usr/local/go/src/testing/testing.go:1118
	// main.main
	// 	_testmain.go:88
	// runtime.main
	// 	/usr/local/go/src/runtime/proc.go:203
	// runtime.goexit
	// 	/usr/local/go/src/runtime/asm_amd64.s:1357
}

func ExampleError_unwrap_with_error() {
	err := errors.New("xerror")
	werr := xerror.Wrap(err, "wrap message")
	fmt.Printf("%+v\n", werr.Unwrap())

	// Output:
	// xerror
	// wrap message
	// github.com/save95/xerror.Wrap
	//	/home/user/dev/xerror/wrap.go:25
	//github.com/save95/xerror/examples.ExampleError_unwrap_with_error
	//	/home/user/dev/xerror/examples/show_stack_test.go:62
	//testing.runExample
	//	/usr/local/go/src/testing/run_example.go:62
	//testing.runExamples
	//	/usr/local/go/src/testing/example.go:44
	//testing.(*M).Run
	//	/usr/local/go/src/testing/testing.go:1118
	//main.main
	//	_testmain.go:88
	//runtime.main
	//	/usr/local/go/src/runtime/proc.go:203
	//runtime.goexit
	//	/usr/local/go/src/runtime/asm_amd64.s:1357
}

func ExampleError_with_fields() {
	err := xerror.New("xerror").WithFields(map[string]interface{}{
		"foo": "bar",
		"bar": 0.79,
		"struct": struct {
			Foo string
		}{
			Foo: "bar",
		},
	})
	fmt.Printf("%+v\n", err)

	// Output:
	// fields: [{"bar":0.79,"foo":"bar","struct":{"Foo":"bar"}}]
	//xerror
	//github.com/save95/xerror.New
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/func.go:15
	//github.com/save95/xerror/examples.ExampleError_with_fields
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/examples/show_stack_test.go:103
	//testing.runExample
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/run_example.go:64
	//testing.runExamples
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/example.go:44
	//testing.(*M).Run
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/testing.go:1505
	//main.main
	//	_testmain.go:91
	//runtime.main
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/proc.go:255
	//runtime.goexit
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/asm_amd64.s:1581
}

func ExampleError_with_fields_wrap() {
	err := xerror.New("xerror").WithFields(map[string]interface{}{
		"foo": "bar",
		"bar": 0.79,
		"struct": struct {
			Foo string
		}{
			Foo: "bar",
		},
	})
	werr := xerror.Wrap(err, "wrap message")
	fmt.Printf("%+v\n", werr)

	// Output:
	// fields: [{"bar":0.79,"foo":"bar","struct":{"Foo":"bar"}}]
	//xerror
	//github.com/save95/xerror.New
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/func.go:15
	//github.com/save95/xerror/examples.ExampleError_with_fields_wrap
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/examples/show_stack_test.go:136
	//testing.runExample
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/run_example.go:64
	//testing.runExamples
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/example.go:44
	//testing.(*M).Run
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/testing.go:1505
	//main.main
	//	_testmain.go:93
	//runtime.main
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/proc.go:255
	//runtime.goexit
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/asm_amd64.s:1581
	//wrap message
	//github.com/save95/xerror.Wrap
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/wrap.go:23
	//github.com/save95/xerror/examples.ExampleError_with_fields_wrap
	//	/Users/royee/Develop/PoeticalSoft/save95/xerror/examples/show_stack_test.go:145
	//testing.runExample
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/run_example.go:64
	//testing.runExamples
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/example.go:44
	//testing.(*M).Run
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/testing/testing.go:1505
	//main.main
	//	_testmain.go:93
	//runtime.main
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/proc.go:255
	//runtime.goexit
	//	/usr/local/Cellar/go@1.17/1.17.13/libexec/src/runtime/asm_amd64.s:1581
}
