package log

import "github.com/sirupsen/logrus"

type LogRusEntry interface {
	Data() logrus.Fields
	Info(args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) LogRusEntry
}

type logRusEntry struct {
	entry *logrus.Entry
}

func (l logRusEntry) Error(args ...interface{}) {
	l.entry.Log(logrus.ErrorLevel, args...)
}

func (l logRusEntry) WithField(key string, value interface{}) LogRusEntry {
	l.entry.WithField(key, value)
	return l
}

func (l logRusEntry) Info(args ...interface{}) {
	l.entry.Log(logrus.InfoLevel, args...)
}

func (l logRusEntry) Data() logrus.Fields {
	return l.entry.Data
}
