package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Level represents the level of log.
type Level byte

const (
	// LevelOff represents turn off all logs.
	LevelOff Level = iota
	// LevelPanic turns on: panic logs only.
	LevelPanic
	// LevelFatal turns on: panic & fatal logs.
	LevelFatal
	// LevelError turns on: panic & fatal & error logs.
	LevelError
	// LevelWarn turns on: panic & fatal & error & warn logs.
	LevelWarn
	// LevelInfo turns on: panic & fatal & error & warn & info logs.
	LevelInfo
	// LevelDebug turns on: panic & fatal & error & warn & info & debug logs.
	LevelDebug
	// LevelAll turns on all logs.
	LevelAll
)

// Logger represents a level logger.
type Logger struct {
	logger *log.Logger
	level  Level
}

// NewLogger returns a new level logger.
func NewLogger(out io.Writer, prefix string) *Logger {
	return &Logger{
		logger: log.New(out, prefix, log.LstdFlags),
		level:  LevelAll,
	}
}

// Level returns the level of logger.
func (l *Logger) Level() Level {
	return l.level
}

// SetLevel sets the level of logger.
func (l *Logger) SetLevel(v Level) {
	l.level = v
}

// Prefix returns the prefix of logger.
func (l *Logger) Prefix() string {
	return l.logger.Prefix()
}

// SetPrefix sets the prefix of logger.
func (l *Logger) SetPrefix(v string) {
	l.logger.SetPrefix(v)
}

// SetOutput sets the output of logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// Panic outputs panic log followed by a call to panic.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Panic(args ...interface{}) {
	if l.level >= LevelPanic {
		l.logger.Panic(args...)
	} else {
		panic(fmt.Sprint(args...))
	}
}

// Panicf outputs panic log followed by a call to panic.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.level >= LevelPanic {
		l.logger.Panicf(format, args...)
	} else {
		panic(fmt.Sprintf(format, args...))
	}
}

// Panicln outputs panic log followed by a call to panic.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Panicln(args ...interface{}) {
	if l.level >= LevelPanic {
		l.logger.Panicln(args...)
	} else {
		panic(fmt.Sprintln(args...))
	}
}

// Fatal outputs fatal log followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Fatal(args ...interface{}) {
	if l.level >= LevelFatal {
		l.logger.Fatal(args...)
	} else {
		os.Exit(1)
	}
}

// Fatalf outputs fatal log followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.level >= LevelFatal {
		l.logger.Fatalf(format, args...)
	} else {
		os.Exit(1)
	}
}

// Fatalln outputs fatal log followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Fatalln(args ...interface{}) {
	if l.level >= LevelFatal {
		l.logger.Fatalln(args...)
	} else {
		os.Exit(1)
	}
}

// Error outputs error log.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {
	if l.level >= LevelError {
		l.logger.Print(args...)
	}
}

// Errorf outputs error log.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level >= LevelError {
		l.logger.Printf(format, args...)
	}
}

// Errorln outputs error log.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Errorln(args ...interface{}) {
	if l.level >= LevelError {
		l.logger.Println(args...)
	}
}

// Warn outputs warn log.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warn(args ...interface{}) {
	if l.level >= LevelWarn {
		l.logger.Print(args...)
	}
}

// Warnf outputs warn log.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level >= LevelWarn {
		l.logger.Printf(format, args...)
	}
}

// Warnln outputs warn log.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warnln(args ...interface{}) {
	if l.level >= LevelWarn {
		l.logger.Println(args...)
	}
}

// Info outputs info log.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {
	if l.level >= LevelInfo {
		l.logger.Print(args...)
	}
}

// Infof outputs info log.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= LevelInfo {
		l.logger.Printf(format, args...)
	}
}

// Infoln outputs info log.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Infoln(args ...interface{}) {
	if l.level >= LevelInfo {
		l.logger.Println(args...)
	}
}

// Debug outputs debug log.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(args ...interface{}) {
	if l.level >= LevelDebug {
		l.logger.Print(args...)
	}
}

// Debugf outputs debug log.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= LevelDebug {
		l.logger.Printf(format, args...)
	}
}

// Debugln outputs debug log.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debugln(args ...interface{}) {
	if l.level >= LevelDebug {
		l.logger.Println(args...)
	}
}

var (
	// DefaultLogger represents the default logger.
	DefaultLogger = NewLogger(os.Stdout, "")
)

// Panic is short for DefaultLogger.Panic.
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// Panicf is short for DefaultLogger.Panicf.
func Panicf(fmt string, args ...interface{}) {
	DefaultLogger.Panicf(fmt, args...)
}

// Panicln is short for DefaultLogger.Panicln.
func Panicln(args ...interface{}) {
	DefaultLogger.Panicln(args...)
}

// Fatal is short for DefaultLogger.Fatal.
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Fatalf is short for DefaultLogger.Fatalf.
func Fatalf(fmt string, args ...interface{}) {
	DefaultLogger.Fatalf(fmt, args...)
}

// Fatalln is short for DefaultLogger.Fatalln.
func Fatalln(args ...interface{}) {
	DefaultLogger.Fatalln(args...)
}

// Error is short for DefaultLogger.Error.
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Errorf is short for DefaultLogger.Errorf.
func Errorf(fmt string, args ...interface{}) {
	DefaultLogger.Errorf(fmt, args...)
}

// Errorln is short for DefaultLogger.Errorln.
func Errorln(args ...interface{}) {
	DefaultLogger.Errorln(args...)
}

// Warn is short for DefaultLogger.Warn.
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Warnf is short for DefaultLogger.Warnf.
func Warnf(fmt string, args ...interface{}) {
	DefaultLogger.Warnf(fmt, args...)
}

// Warnln is short for DefaultLogger.Warnln.
func Warnln(args ...interface{}) {
	DefaultLogger.Warnln(args...)
}

// Info is short for DefaultLogger.Info.
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Infof is short for DefaultLogger.Infof.
func Infof(fmt string, args ...interface{}) {
	DefaultLogger.Infof(fmt, args...)
}

// Infoln is short for DefaultLogger.Infoln.
func Infoln(args ...interface{}) {
	DefaultLogger.Infoln(args...)
}

// Debug is short for DefaultLogger.Debug.
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Debugf is short for DefaultLogger.Debugf.
func Debugf(fmt string, args ...interface{}) {
	DefaultLogger.Debugf(fmt, args...)
}

// Debugln is short for DefaultLogger.Debugln.
func Debugln(args ...interface{}) {
	DefaultLogger.Debugln(args...)
}
