package loghelper

import (
	"context"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// app logger
var appLogger *zap.SugaredLogger

func initHelper(ctx context.Context) {
	// writerSyncer := getLogWriter()
	_, consoleEncoder := getEncoder()
	// fileEncoder, consoleEncoder := getEncoder()
	defaultLogLevel := zap.DebugLevel
	core := zapcore.NewTee(
		// zapcore.NewCore(fileEncoder, writerSyncer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	backgroundCore, err := nrzap.WrapTransactionCore(core, newrelic.FromContext(ctx))

	if err != nil && err != nrzap.ErrNilTxn {
		// panic(err)
		log.Printf("Cannot connect to new relic zap:%v\n", err)
	}

	appLogger := zap.New(backgroundCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	// logger := zap.New(ecszap.WrapCore(core), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	// appLogger = logger.Sugar()
	defer appLogger.Sync()
}

func getEncoder() (zapcore.Encoder, zapcore.Encoder) {
	// encoderConfig := ecszap.ECSCompatibleEncoderConfig(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	return fileEncoder, consoleEncoder
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./storage/logs/app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Debug(args...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Debugf(template, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Info(args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Infof(template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Warn(args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Warnf(template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Error(args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Errorf(template, args...)
}

func DPanic(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).DPanic(args...)
}

func DPanicf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).DPanicf(template, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Panic(args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Panicf(template, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Fatal(args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	initHelper(ctx)
	setKeyLogger(ctx, appLogger).Fatalf(template, args...)
}
func setKeyLogger(ctx context.Context, log *zap.SugaredLogger) *zap.SugaredLogger {
	return log
	//// initHelper(ctx)
	//// xTraceID := ctx.Value(XTRACEID).(string)
	// xName := "log" + xTraceID
	// if ctx.Value(XAPPNAME) != nil {
	// 	xName = ctx.Value(XAPPNAME).(string)
	// }
	// return log.With(zap.String("trace-id", xTraceID)).Named(xName)
}
