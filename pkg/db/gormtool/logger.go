package gormtool

import (
	"context"
	"errors"
	"fmt"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log"
	"os"
	"time"
)

var DefaultLogger = NewLogger(log.New(os.Stdout, "", log.LstdFlags), gormlog.Config{
	SlowThreshold: 200 * time.Millisecond,
	LogLevel:      gormlog.Warn,
	Colorful:      false,
})

func NewLogger(writer gormlog.Writer, config gormlog.Config) gormlog.Interface {
	var (
		infoStr      = "%s [info] %s"
		warnStr      = "%s [warn] %s"
		errStr       = "%s [error] %s"
		traceStr     = "%s [%.3fms] [rows:%v] %s %s"
		traceWarnStr = "%s %s [%.3fms] [rows:%v] %s %s"
		traceErrStr  = "%s %s [%.3fms] [rows:%v] %s %s"
	)

	return &logger{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type logger struct {
	gormlog.Writer
	gormlog.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *logger) LogMode(level gormlog.LogLevel) gormlog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlog.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum(), l.tracingInfo(ctx)}, data...)...)
	}
}

// Warn print warn messages
func (l logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlog.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum(), l.tracingInfo(ctx)}, data...)...)
	}
}

// Error print error messages
func (l logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlog.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum(), l.tracingInfo(ctx)}, data...)...)
	}
}

// Trace print sql message
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.LogLevel <= gormlog.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlog.Error && (!errors.Is(err, gormlog.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", l.tracingInfo(ctx), sql)
		} else {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, l.tracingInfo(ctx), sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlog.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", l.tracingInfo(ctx), sql)
		} else {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, l.tracingInfo(ctx), sql)
		}
	case l.LogLevel == gormlog.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", l.tracingInfo(ctx), sql)
		} else {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, l.tracingInfo(ctx), sql)
		}
	}
}

func (l *logger) tracingInfo(ctx context.Context) string {
	//span := tracing.SpanFromContext(ctx)
	//if span != nil {
	//	return fmt.Sprintf("[trace:req:%s,account:%s]", span.RequestID(), span.AuthAccountID())
	//}
	return "-"
}
