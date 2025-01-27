package stdfn

import (
	"github.com/go-stdlog/stdlog"
)

type ConsumerFunc func(level stdlog.Level, name string, err error, msg string, keyValues []any)

type fnLogger struct {
	target     ConsumerFunc
	name       string
	level      stdlog.Level
	baseFields []any
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

func (f fnLogger) WithFields(keysAndValues ...any) stdlog.Logger {
	assertKvs("WithFields", keysAndValues)
	f.baseFields = append(f.baseFields, keysAndValues...)
	return f
}

func assertKvs(method string, kvs []any) {
	if l := len(kvs); l != 0 && l%2 != 0 {
		panic("uneven number of key-value pairs passed to " + method)
	}
}

func (f fnLogger) Debug(msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelDebug.String(), keysAndValues)
	f.target(stdlog.LevelDebug, f.name, nil, msg, append(f.baseFields, keysAndValues...))
}

func (f fnLogger) Info(msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelInfo.String(), keysAndValues)
	f.target(stdlog.LevelInfo, f.name, nil, msg, append(f.baseFields, keysAndValues...))
}

func (f fnLogger) Warning(msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelWarning.String(), keysAndValues)
	f.target(stdlog.LevelWarning, f.name, nil, msg, append(f.baseFields, keysAndValues...))
}

func (f fnLogger) Error(err error, msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelError.String(), keysAndValues)
	f.target(stdlog.LevelError, f.name, err, msg, append(f.baseFields, keysAndValues...))
}

func (f fnLogger) Fatal(msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelFatal.String(), keysAndValues)
	f.target(stdlog.LevelFatal, f.name, nil, msg, append(f.baseFields, keysAndValues...))
}

func (f fnLogger) FatalError(err error, msg string, keysAndValues ...any) {
	assertKvs(stdlog.LevelFatal.String(), keysAndValues)
	f.target(stdlog.LevelFatal, f.name, err, msg, append(f.baseFields, keysAndValues...))
}

func New(consumer ConsumerFunc) stdlog.Logger {
	return &fnLogger{target: consumer}
}
