# xerror

错误码处理包：包含基础的错误处理

## Usage

### 基本用法

```go
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

func ExampleWithXCode() {
	err := xerror.WithXCode(xcode.InternalServerError)
	fmt.Println(err)

	// Output:
	// 内部服务错误
}
```

更多用法请访问 [错误码基本使用场景](examples/base_usage_test.go)

### 进阶用法

  - [客户端消息](examples/client_message_test.go)
  - [错误断言](examples/compare_test.go)
  - [加载/覆写错误码](examples/load_xcode_test.go)
  - [打印错误栈](examples/show_stack_test.go)
  - [携带自定义参数的错误](examples/error_with_filed_test.go)
  - [包装错误（wrap）](examples/wrap_test.go)

# License

[MIT License](LICENSE)
