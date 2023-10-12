package xerror_test

import (
	"testing"

	"github.com/save95/xerror"
	"github.com/save95/xerror/ecode"
	"github.com/save95/xerror/xcode"
)

func TestNew(t *testing.T) {
	t.Log(xcode.New(500))
	t.Log(xerror.New("this is new error"))

	err := xerror.WithXCode(xcode.InternalServerError)
	t.Log(err, " / ", err.ToMessage(&ecode.Config{}))
}

func TestNewWithCode(t *testing.T) {
	err := xerror.WithCode(100001, "木有登陆")
	t.Log(err)
	t.Log(err.ToMessage(nil))
}

func TestNewWithXCode(t *testing.T) {
	err := xerror.WithXCode(xcode.InternalServerError)
	t.Log(err)
	t.Log(err.ToMessage(nil))
}

func TestWithXCodeMessage(t *testing.T) {
	err := xerror.WithXCodeMessage(xcode.InternalServerError, "变更消息")
	t.Log(err)
	t.Log(err.ToMessage(nil))
}

func TestParsePayload(t *testing.T) {
	err := xerror.WithXCodeMessage(xcode.InternalServerError, "变更消息")
	t.Log(xerror.ParsePayload(err))

	err2 := xerror.Wrap(err, "错误2").
		WithFields("abc").
		WithFields(map[string]interface{}{
			"k": "v",
			"a": 1,
		})
	t.Log(xerror.ParsePayload(err2))
}

func TestStackTraceString(t *testing.T) {
	err := xerror.WithXCodeMessage(xcode.InternalServerError, "变更消息")
	t.Log(xerror.FormatStackTrace(err))

	err2 := xerror.Wrap(err, "错误2").
		WithFields("abc").
		WithFields(map[string]interface{}{
			"k": "v",
			"a": 1,
		})
	t.Log(xerror.FormatStackTrace(err2))
}
