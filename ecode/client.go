package ecode

// 端口
type Client uint

const (
	ClientDefault             Client = iota // 默认
	ClientWeb                               // Web 端口
	ClientH5                                // H5 端口
	ClientApp                               // APP 端口（不分设备类型）
	ClientIOS                               // IOS 设备端口
	ClientAndroid                           // 安卓设备端口
	ClientWechatPublicAccount               // 微信公众号端口
	ClientWechatMiniProgram                 // 微信小程序端口
)
