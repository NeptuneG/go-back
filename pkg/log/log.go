package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel
	WarnLevel   Level = zap.WarnLevel
	ErrorLevel  Level = zap.ErrorLevel
	DPanicLevel Level = zap.DPanicLevel
	PanicLevel  Level = zap.PanicLevel
	FatalLevel  Level = zap.FatalLevel
	DebugLevel  Level = zap.DebugLevel
)

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.zapLogger.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

var (
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
)

// not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug
}

type Logger struct {
	zapLogger *zap.Logger
	level     Level
}

var std = createStd()

func Default() *Logger {
	return std
}

// not support log rotating
func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("writer is nil")
	}
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		zapLogger: zap.New(core),
		level:     level,
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func createStd() *Logger {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zapLogger, err := config.Build()
	if err != nil {
		panic("failed to initialize zap logger")
	}
	return &Logger{
		zapLogger: zapLogger,
		level:     DebugLevel,
	}
}
