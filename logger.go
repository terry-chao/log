package log

import "sync"

var std = New()

const (
	FmtEmptySeparate = ""
	FmtLineSeparate  = "\n"
)

type logger struct {
	opt       *option
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	logger := &logger{opt: initOptions(opts...)}
	logger.entryPool = &sync.Pool{New: func() interface{} { return entry(logger) }}
	return logger
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}
