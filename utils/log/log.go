package log

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 日志记录器接口
type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warn(message string, args ...interface{})
}

type logger struct {
	log *zap.Logger
}

// NewLogger 创建日志记录器
func NewLogger() Logger {
	// 创建 logs 目录
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0o755)
		if err != nil {
			panic("无法创建 logs 目录: " + err.Error())
		}
	}

	// 获取当前日期作为日志文件名
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := filepath.Join(logDir, logFileName)

	// 创建日志文件输出 writer（带自动切割）
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    100, // 每个文件最大 100MB
		MaxAge:     30,  // 最多保留 30 天
		MaxBackups: 5,   // 最多备份文件数
		Compress:   true,
	})

	// 控制台输出编码器（带颜色）
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色输出
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	// 文件输出编码器（无颜色）
	fileEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 无颜色
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	// 设置日志级别
	level := zapcore.InfoLevel

	// 多输出：控制台 + 文件
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(fileEncoder, fileWriter, level),
	)

	// 构建 logger 实例
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &logger{
		log: log,
	}
}

// customTimeEncoder 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02-15:04:05"))
}

// Info 记录 info 级别日志
func (l *logger) Info(message string, args ...interface{}) {
	l.log.Info(message, zap.Any("args", args))
}

// Error 记录 error 级别日志
func (l *logger) Error(message string, args ...interface{}) {
	l.log.Error(message, zap.Any("args", args))
}

// Debug 记录 debug 级别日志
func (l *logger) Debug(message string, args ...interface{}) {
	l.log.Debug(message, zap.Any("args", args))
}

// Warn 记录 warn 级别日志
func (l *logger) Warn(message string, args ...interface{}) {
	l.log.Warn(message, zap.Any("args", args))
}
