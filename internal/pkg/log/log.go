// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

// Package log is a log package used by dazBlog project
package log

import (
	"context"
	"github.com/Daz-3ux/dBlog/internal/pkg/known"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"sync"
	"time"
)

// Logger define dBlog's logger interface
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

// zapLogger wraps zap.Logger
// the specific implementation of the Logger interface
type zapLogger struct {
	z *zap.Logger
}

// ensure that zapLogger implements the Logger interface
var _ Logger = (*zapLogger)(nil)

var (
	mu sync.Mutex

	// std the global default logger
	std = NewLogger(NewOptions())
)

// Init initializes the global logger with the given options
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// NewLogger creates a new logger with the given options
func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// convert text-based log level, e.g. "info", to the zapcore.Level type
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// if specify an invalid log level, use info level as default
		zapLevel = zapcore.InfoLevel
	}

	// create a default zapcore.EncoderConfig
	encoderConfig := zap.NewProductionEncoderConfig()

	// customize the default zapcore.EncoderConfig
	// customize MessageKey to "message" for a more explicit meaning
	encoderConfig.MessageKey = "message"
	// customize TimeKey to "timestamp" for a more explicit meaning
	encoderConfig.TimeKey = "timestamp"
	// customize Level style to Capital and Color
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// customize Time format for improved readability
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// customize a Duration format for improved precision
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// cfg the config required building a zap.Logger
	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		// specify the default inner error output path
		ErrorOutputPaths: []string{"stderr"},
	}

	// use cfg to build a *zap.Logger
	z, err := cfg.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// OOP: wrap the zap library in a custom struct
	logger := &zapLogger{z: z}

	// redirect the standard library's log to the zap logger
	zap.RedirectStdLog(z)

	return logger
}

// Sync flushes all buffered logs into the disk
// main function should call this function before exit
func Sync() {
	std.Sync()
}

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

// Debugw print debug level log
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

// Infow print info level log
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

// Warnw print warn level log
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw print error level log
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

// Panicw print panic level log
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw print fatal level log
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

// C extracts relevant key-value pairs from the incoming context
// and adds them to the structured logs of the zap.Logger.
func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()

	// Add the request ID to the output fields of the log
	if requestID := ctx.Value(known.XRequestIDKey); requestID != nil {
		lc.z = lc.z.With(zap.Any(known.XRequestIDKey, requestID))
	}

	// Add the username to the output fields of the log
	if userID := ctx.Value(known.XUsernameKey); userID != nil {
		lc.z = lc.z.With(zap.Any(known.XUsernameKey, userID))
	}

	return lc
}

// clone deep copies the zapLogger
// because the log package is called concurrently by multiple request,
// X-Request-ID is protected against contamination
func (l *zapLogger) clone() *zapLogger {
	lc := *l
	return &lc
}
