package ldLoggerKit_test

import (
	"testing"

	"github.com/hkloudou/ldLoggerKit"
)

func Test_log(t *testing.T) {
	var SYSTEM = (&ldLoggerKit.LogBuilder{}).InitLogerFactory("system")
	SYSTEM.Debugln("test")
	SYSTEM.Debugf("%s", "some string")
}
