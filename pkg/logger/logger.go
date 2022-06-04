package logger

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/tool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

func getFileLogWriter(logFilePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10, // 单个文件最大10M
		MaxBackups: 60, // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     1,  // 一天一切割
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func defaultLogWriter() zapcore.WriteSyncer {
	logFilePath := ""

	projectPath := tool.GetProjectAbsPath()

	if os.Getenv(config.ENVkEY) == config.PROD {
		logFilePath = projectPath + "/log/prod.log"
	} else {
		logFilePath = projectPath + "/log/dev.log"
	}

	return getFileLogWriter(logFilePath)
}

func Default() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	fileWriteSyncer := defaultLogWriter()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	)
	logger = zap.New(core)
	return logger
}

func New(logFilePath string) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	file, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	fileWriteSyncer := zapcore.AddSync(file)

	var core zapcore.Core

	core = zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	)

	if os.Getenv(config.ENVkEY) == config.PROD {
		core = zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel)
	}

	return zap.New(core)
}
