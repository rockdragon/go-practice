package main

import (
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
	"github.com/rockdragon/go-practice/router"
	"go.uber.org/zap/zapcore"
	"strings"
	"runtime"
)

func newLoggerConfig(debugLevel bool) (loggerConfig zap.Config) {
	loggerConfig = zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeCaller = CallerEncoder
	loggerConfig.OutputPaths = []string {"./logs/app.log"}
	if debugLevel {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return
}

func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(),
		runtime.FuncForPC(caller.PC).Name()}, ":"))
}

func main() {
	loggerConfig := newLoggerConfig(true)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logger.Info("hello")

	router.Router(gin.Default())
}
