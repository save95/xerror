package xerror

import "github.com/save95/xerror/xcode"

type XLoadCodes interface {
	// 加载自定义 code
	// 该方法可以覆写系统预置的错误码
	LoadCodes(codes ...xcode.XCode)
	// 追加自定义 code
	// 如果已存在，则直接跳过；如果需要覆写，请使用 `LoadCodes`
	AppendCodes(codes ...xcode.XCode)
}
