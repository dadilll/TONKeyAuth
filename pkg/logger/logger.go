package logger

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

type ctxKey string

const (
	LoggerKey   ctxKey = "logger"
	RequestID   ctxKey = "requestID"
	ServiceName        = "service"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Sync() error
}

type logger struct {
	serviceName string
	logger      *zap.Logger
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.log(ctx, msg, fields, "info")
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.log(ctx, msg, fields, "error")
}

func (l *logger) log(ctx context.Context, msg string, fields []zap.Field, level string) {
	if requestID, ok := ctx.Value(RequestID).(string); ok && requestID != "" {
		fields = append(fields, zap.String(string(RequestID), requestID))
	}
	fields = append(fields, zap.String(ServiceName, l.serviceName))

	switch level {
	case "info":
		l.logger.Info(msg, fields...)
	case "error":
		l.logger.Error(msg, fields...)
	}
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}

var (
	mu      sync.Mutex
	loggers = make(map[string]Logger)
)

// New создает новый логгер для сервиса (или возвращает уже созданный)
func New(serviceName string) Logger {
	mu.Lock()
	defer mu.Unlock()

	if log, exists := loggers[serviceName]; exists {
		return log
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	newLogger := &logger{
		serviceName: serviceName,
		logger:      zapLogger,
	}
	loggers[serviceName] = newLogger
	return newLogger
}

// WithLogger добавляет логгер в контекст
func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, log)
}

// GetLoggerFromCtx получает логгер из контекста или возвращает дефолтный
func GetLoggerFromCtx(ctx context.Context) Logger {
	if logger, ok := ctx.Value(LoggerKey).(Logger); ok {
		return logger
	}
	return New("default")
}
