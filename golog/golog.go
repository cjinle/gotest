// Package golog 日志操作，按天分割
package golog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	Error int8 = iota
	Warn
	Info
	Debug
)

type LogConf struct {
	Path   string // 日志目录
	Level  int8   // 日志级别 3:DEBUG 2:INFO 1:WARN 0:ERROR
	Prefix string // 日志文件名前缀(一般用于区分同一日志目录下多类日志文件)
}

type Logger struct {
	conf     *LogConf
	logger   *log.Logger
	fileMu   *sync.Mutex // 控制日志文件切换
	file     *os.File    // 日志文件句柄
	fileTime time.Time   // 日志文件创建时间
}

func NewLogger(conf *LogConf) *Logger {
	logger := &Logger{
		conf:     conf,
		fileMu:   new(sync.Mutex),
		fileTime: time.Now(),
	}
	if logger.conf.Path != "" {
		if err := os.MkdirAll(logger.conf.Path, 0755); err != nil {
			log.Fatalf("NewLogger|Make LogPath=%s|%s", logger.conf.Path, err)
		}
	}
	y, m, d := logger.fileTime.Date()
	filename := fmt.Sprintf("%s/%s%d_%d%02d%02d.log", logger.conf.Path,
		logger.conf.Prefix, os.Getpid(), y, int(m), d)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("NewLogger|Create LogFile=%s|%s", filename, err)
	}
	logger.file = file
	logger.logger = log.New(logger.file, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.Printf("LogFile=%s", filename)
	return logger
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) Printf(level int8, format string, a ...interface{}) {
	if level > l.conf.Level {
		return
	}
	var logFmt string
	switch level {
	case Error:
		logFmt = "[ERROR]" + format
	case Warn:
		logFmt = "[WARN]" + format
	case Info:
		logFmt = "[INFO]" + format
	case Debug:
		logFmt = "[DEBUG]" + format
	}
	now := time.Now()
	y, m, d := now.Date()

	l.fileMu.Lock()
	ly, lm, ld := l.fileTime.Date()
	if ld != d || lm != m || ly != y { // 切换新文件
		filename := fmt.Sprintf("%s/%s%d_%d%02d%02d.log", l.conf.Path, l.conf.Prefix, os.Getpid(), y, int(m), d)
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			l.logger.Printf("Printf|Create LogFile=%s|%s", filename, err)
			l.fileMu.Unlock()
			return
		}
		l.fileTime = now
		l.file.Close()
		l.file = file
		l.logger.SetOutput(l.file)
	}
	l.fileMu.Unlock()

	l.logger.Output(2, fmt.Sprintf(logFmt, a...))
}
