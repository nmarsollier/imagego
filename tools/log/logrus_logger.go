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
	l.entry.Error(args...)
}

func (l logRusEntry) WithField(key string, value interface{}) LogRusEntry {
	return logRusEntry{
		entry: l.entry.WithField(key, value),
	}
}

func (l logRusEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l logRusEntry) Data() logrus.Fields {
	return l.entry.Data
}
