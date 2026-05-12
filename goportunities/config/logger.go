package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

// create Non-Formatted logs

func (logger *Logger) Debug(value ...any) {
	logger.debug.Println(value...)
}

func (logger *Logger) Info(value ...any) {
	logger.info.Println(value...)
}

func (logger *Logger) Warning(value ...any) {
	logger.warning.Println(value...)
}

func (logger *Logger) Error(value ...any) {
	logger.error.Println(value...)
}

// create Formatted logs

func (logger *Logger) Debugf(format string, value ...any) {
	logger.debug.Printf(format, value...)
}

func (logger *Logger) Infof(format string, value ...any) {
	logger.info.Printf(format, value...)
}

func (logger *Logger) Warningf(format string, value ...any) {
	logger.warning.Printf(format, value...)
}

func (logger *Logger) Errorf(format string, value ...any) {
	logger.error.Printf(format, value...)
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, "> DEBUG: ", logger.Flags()),
		info:    log.New(writer, "> INFO: ", logger.Flags()),
		warning: log.New(writer, "> WARNING: ", logger.Flags()),
		error:   log.New(writer, "> ERROR: ", logger.Flags()),
		writer:  writer,
	}
}
