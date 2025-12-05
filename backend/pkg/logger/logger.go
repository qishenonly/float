package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

// Init 初始化日志
func Init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(config)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()
}

// Debug 日志
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Debugf 格式化Debug日志
func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

// Info 日志
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof 格式化Info日志
func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

// Warn 日志
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf 格式化Warn日志
func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

// Error 日志
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf 格式化Error日志
func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

// Fatal 日志
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf 格式化Fatal日志
func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
