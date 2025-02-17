package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

var logger *logrus.Logger

func init() {
	logger = GetLogger()
}

func GetLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}
	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{})
	return lg
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return logger.WithError(err)
}

func WithContext(ctx context.Context) *logrus.Entry {
	return logger.WithContext(ctx)
}

func WithTime(t time.Time) *logrus.Entry {
	return logger.WithTime(t)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	logger.Logf(level, format, args...)
}

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Log(level logrus.Level, args ...interface{}) {
	logger.Log(level, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Logln(level logrus.Level, args ...interface{}) {
	logger.Logln(level, args...)
}

func Traceln(args ...interface{}) {
	logger.Traceln(args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	logger.Warnln(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}
