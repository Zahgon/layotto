/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"sync"

	"mosn.io/pkg/log"
)

const (
	// TraceLevel is for logging verbose message with a set of methods and properties to help track code execution.
	TraceLevel LogLevel = "trace"
	// DebugLevel has verbose message.
	DebugLevel LogLevel = "debug"
	// InfoLevel is default log level.
	InfoLevel LogLevel = "info"
	// WarnLevel is for logging messages about possible issues.
	WarnLevel LogLevel = "warn"
	// ErrorLevel is for logging errors.
	ErrorLevel LogLevel = "error"
	// FatalLevel is for logging fatal messages.
	FatalLevel LogLevel = "fatal"

	// UndefinedLevel is for undefined log level.
	UndefinedLevel LogLevel = "undefined"

	logKeyDebug    = "debug"
	logKeyAccess   = "access"
	logKeyError    = "error"
	fileNameDebug  = "layotto.debug.log"
	fileNameAccess = "layotto.access.log"
	fileNameError  = "layotto.error.log"
)

var (
	loggerListeners    sync.Map
	defaultLoggerLevel = InfoLevel
	defaultLogFilePath = "./"
)

// LogLevel is Logger Level type.
type LogLevel string

// ComponentLoggerListener is the interface for setting log config.
type ComponentLoggerListener interface {
	OnLogLevelChanged(outputLevel LogLevel)
}

// RegisterComponentLoggerListener registers a logger for a component logger listener.
func RegisterComponentLoggerListener(componentName string, logger ComponentLoggerListener) {
	_ = "STUB: not implemented"
	return
}

// SetComponentLoggerLevel sets the log level for a component.
func SetComponentLoggerLevel(componentName string, level string) { _ = "STUB: not implemented"; return }

// SetDefaultLoggerLevel sets the default log output level.
func SetDefaultLoggerLevel(level string) { _ = "STUB: not implemented"; return }

// SetDefaultLoggerFilePath sets the default log file path.
func SetDefaultLoggerFilePath(filePath string) { _ = "STUB: not implemented"; return }

// layottoLogger is the implementation for layotto.
type layottoLogger struct {
	// name is the name of logger that is published to log as a component.
	name string

	logLevel LogLevel

	loggers map[string]log.ErrorLogger
}

// Logger api for logging.
type Logger interface {
	// Trace logs a message at level Trace.
	Trace(args ...interface{})
	// Tracef logs a message at level Trace.
	Tracef(format string, args ...interface{})
	// Debug logs a message at level Debug.
	Debug(args ...interface{})
	// Debugf logs a message at level Debug.
	Debugf(format string, args ...interface{})
	// Info logs a message at level Info.
	Info(args ...interface{})
	// Infof logs a message at level Info.
	Infof(format string, args ...interface{})
	// Warn logs a message at level Warn.
	Warn(args ...interface{})
	// Warnf logs a message at level Warn.
	Warnf(format string, args ...interface{})
	// Error logs a message at level Error.
	Error(args ...interface{})
	// Errorf logs a message at level Error.
	Errorf(format string, args ...interface{})
	// Fatal logs a message at level Fatal.
	Fatal(args ...interface{})
	// Fatalf logs a message at level Fatal.
	Fatalf(format string, args ...interface{})
	// SetLogLevel sets the log output level
	SetLogLevel(outputLevel LogLevel)
	// GetLogLevel get the log output level
	GetLogLevel() LogLevel
}

// toLogLevel converts to LogLevel.
func toLogLevel(level string) LogLevel { _ = "STUB: not implemented"; return *new(LogLevel) }

// unsupported log level

// ToLogPriority converts to Logger priority.
func ToLogPriority(level LogLevel) int { _ = "STUB: not implemented"; return 0 }

// NewLayottoLogger creates new Logger instance.
func NewLayottoLogger(name string) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// Tracef logs a message at level Trace.
func (l *layottoLogger) Tracef(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Trace logs a message at level Trace.
func (l *layottoLogger) Trace(args ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf logs a message at level Debug.
func (l *layottoLogger) Debugf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Debug logs a message at level Debug.
func (l *layottoLogger) Debug(args ...interface{}) { _ = "STUB: not implemented"; return }

// Infof logs a message at level Info.
func (l *layottoLogger) Infof(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Info logs a message at level Info.
func (l *layottoLogger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf logs a message at level Warn.
func (l *layottoLogger) Warnf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Warn logs a message at level Warn.
func (l *layottoLogger) Warn(args ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf logs a message at level Error.
func (l *layottoLogger) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Error logs a message at level Error.
func (l *layottoLogger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf logs a message at level Fatal.
func (l *layottoLogger) Fatalf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Fatal logs a message at level Fatal.
func (l *layottoLogger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

// GetLogLevel gets the log output level.
func (l *layottoLogger) GetLogLevel() LogLevel {
	_ = "STUB: not implemented"

	// toMosnLoggerLevel converts to logrus.Level.
	return *new(LogLevel)
}

func toMosnLoggerLevel(lvl LogLevel) log.Level {
	_ = "STUB: not implemented"
	// ignore error because it will never happen
	return *new(log.Level)
}

// parseLevel takes a string level and returns the Mosn logger level constant.
func parseLevel(lvl string) (log.Level, error) {
	_ = "STUB: not implemented"
	return *new(log.Level), nil
}

// SetLogLevel sets log output level.
func (l *layottoLogger) SetLogLevel(outputLevel LogLevel) { _ = "STUB: not implemented"; return }
