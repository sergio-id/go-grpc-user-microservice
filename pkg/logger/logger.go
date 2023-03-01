package logger

import (
	"github.com/sergio-id/go-grpc-user-microservice/pkg/constants"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface
type Logger interface {
	InitLogger()
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Printf(template string, args ...interface{})
	Fatal(args ...interface{})
	Named(name string)
	GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error)
	GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error)
}

// AppLogger is a logger
type appLogger struct {
	level       string
	console     bool
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

// NewAppLogger creates new logger
func NewAppLogger(cfg Config) *appLogger {
	return &appLogger{level: cfg.LogLevel, console: cfg.Console}
}

// InitLogger initializes logger
func (l *appLogger) InitLogger() {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.NameKey = "[NAME]"
	encoderConfig.TimeKey = "[TIME]"
	encoderConfig.LevelKey = "[LEVEL]"
	encoderConfig.CallerKey = "[CALLER]"
	encoderConfig.FunctionKey = "[FUNC]"
	encoderConfig.MessageKey = "[MESSAGE]"
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	var encoder zapcore.Encoder
	if l.console {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoderConfig.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// New core with encoder and writer
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(convertLogLevel(l.level)))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger, l.sugarLogger = logger, logger.Sugar()
}

func convertLogLevel(level string) zapcore.Level {
	var l zapcore.Level

	switch strings.ToLower(level) {
	case "debug":
		l = zapcore.DebugLevel
	case "info":
		l = zapcore.InfoLevel
	case "warn":
		l = zapcore.WarnLevel
	case "error":
		l = zapcore.ErrorLevel
	case "dpanic":
		l = zapcore.DPanicLevel
	case "panic":
		l = zapcore.PanicLevel
	case "fatal":
		l = zapcore.FatalLevel
	default:
		l = zapcore.DebugLevel
	}

	return l
}

func (l *appLogger) Named(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

func (l *appLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *appLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *appLogger) Printf(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *appLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *appLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func (l *appLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *appLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit(1).
func (l *appLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *appLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Info(
		constants.GRPC,
		zap.String(constants.METHOD, method),
		zap.Duration(constants.TIME, time),
		zap.Any(constants.METADATA, metaData),
		zap.Any(constants.ERROR, err),
	)
}

func (l *appLogger) GrpcClientInterceptorLogger(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Info(
		constants.GRPC,
		zap.String(constants.METHOD, method),
		zap.Any(constants.REQUEST, req),
		zap.Any(constants.REPLY, reply),
		zap.Duration(constants.TIME, time),
		zap.Any(constants.METADATA, metaData),
		zap.Any(constants.ERROR, err),
	)
}
