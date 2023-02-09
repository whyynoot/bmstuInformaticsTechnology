package logging

import (
	"go.uber.org/zap/zapcore"
)

// Field is weired, but need to be here because of zap limitations, easy to change for any logger you want to implement
type Field = zapcore.Field

// LoggerInterface is an interface to interact with logger
type LoggerInterface interface {
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}
