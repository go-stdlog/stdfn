package stdfn

import "github.com/go-stdlog/stdlog"

type ConsumerFunc func(level stdlog.Level, name string, err error, msg string, keyValues []any)

type fnLogger struct {
	target ConsumerFunc
	name   string
	level  stdlog.Level
}

func (f fnLogger) Named(name string) stdlog.Logger {
	if f.name != "" {
		f.name = f.name + "." + name
	} else {
		f.name = name
	}
	return f
}

func (f fnLogger) SetLevel(level stdlog.Level) {
	f.level = level
}

func (f fnLogger) Leveled(level stdlog.Level) stdlog.Logger {
	f.level = level
	return f
}

func (f fnLogger) Debug(msg string, keysAndValues ...any) {
	f.target(stdlog.LevelDebug, f.name, nil, msg, keysAndValues)
}

func (f fnLogger) Info(msg string, keysAndValues ...any) {
	f.target(stdlog.LevelInfo, f.name, nil, msg, keysAndValues)
}

func (f fnLogger) Warning(msg string, keysAndValues ...any) {
	f.target(stdlog.LevelWarning, f.name, nil, msg, keysAndValues)
}

func (f fnLogger) Error(err error, msg string, keysAndValues ...any) {
	f.target(stdlog.LevelError, f.name, err, msg, keysAndValues)
}

func (f fnLogger) Fatal(msg string, keysAndValues ...any) {
	f.target(stdlog.LevelFatal, f.name, nil, msg, keysAndValues)
}

func (f fnLogger) FatalError(err error, msg string, keysAndValues ...any) {
	f.target(stdlog.LevelFatal, f.name, err, msg, keysAndValues)
}

func New(consumer ConsumerFunc) stdlog.Logger {
	return &fnLogger{target: consumer}
}
