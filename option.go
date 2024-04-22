package log

import (
	"io"
	"os"
)

type option struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*option)

func initOptions(opts ...Option) (o *option) {
	o = &option{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

func WithLevel(level Level) Option {
	return func(o *option) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *option) {
		o.stdLevel = level
	}
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}
