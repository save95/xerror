package xerror

import "github.com/save95/xerror/xcode"

// LoadCodes 加在外部 XCode 定义
func LoadCodes(codes ...xcode.XCode) {
	xcode.Repository().LoadCodes(codes...)
}

// AppendCodes 追加外部 XCode 定义
func AppendCodes(codes ...xcode.XCode) {
	xcode.Repository().AppendCodes(codes...)
}
