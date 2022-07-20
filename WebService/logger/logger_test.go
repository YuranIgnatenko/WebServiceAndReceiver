package logger

import (
	"go/build"
	"testing"
)

func Test_Logger(t *testing.T) {
	NAMELOGFILE = "test.log"
	PATH = build.Default.GOPATH + "src/WebService/logger/"
	logger := NewLogger()
	logger.INFO("test info")
	logger.WARN("test warn")
	logger.ERROR("test error")
}
