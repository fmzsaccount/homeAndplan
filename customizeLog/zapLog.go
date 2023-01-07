package customizeLog

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gy_home/internal/conf"
	"os"
	"time"
)

//var ProviderSet = wire.NewSet(NewZapLogger)
var _ log.Logger = (*ZapLogger)(nil)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

type Ext struct {
	Ip     string `json:"ip,omitempty"`
	Mac    string `json:"mac,omitempty"`
	Name   string `json:"name,omitempty"`
	Module string `json:"module,omitempty"`
	Action string `json:"action,omitempty"`
}

type LogToKafka struct {
	Brokers       []string
	Topic         string
	WriterTimeout time.Duration
	writer        *kafka.Writer
	Async         bool
}

func (ltk *LogToKafka) Write(buf []byte) (n int, err error) {
	ltk.writer = &kafka.Writer{
		Addr:         kafka.TCP(ltk.Brokers...),
		Topic:        ltk.Topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: ltk.WriterTimeout,
		Async:        ltk.Async,
	}
	kbuf := append([]byte(nil), buf...)
	err = ltk.writer.WriteMessages(context.Background(), kafka.Message{Value: kbuf})
	return
}
// Logger 配置zap日志,将zap日志库引入
func Logger(c *conf.Log) log.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:   "tm",
		LevelKey:  "lvl",
		NameKey:   "zap",
		CallerKey: "src",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var defaultFields []zap.Field
	defaultFields = append(defaultFields,
		zap.Any("appid", c.Sys.Appid),
		zap.Any("env", c.Sys.Env), )
	return NewZapLogger(
		c,
		encoder,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.Development(),
		zap.Fields(defaultFields...),
	)
}


// NewZapLogger return a zap logger.
func NewZapLogger(c *conf.Log,encodercfg zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	kl := &LogToKafka{
		Brokers: c.Kafka.Broker,
		Topic: c.Kafka.Topic,
		WriterTimeout: c.Kafka.WriteTimeout.AsDuration(),
		Async: c.Kafka.Async,
	}
	level.SetLevel(zap.InfoLevel)
	var (
		Core zapcore.Core
	)

	Core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encodercfg),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(kl)),
		level,
		)
	zapLogger := zap.New(Core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}

// Server is an server logging middleware.
func Server(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			//var (
			//	code      int32
			//	reason    string
			//	kind      string
			//	operation string
			//)
			//startTime := time.Now()
			//if info, ok := transport.FromServerContext(ctx); ok {
			//	kind = info.Kind().String()
			//	operation = info.Operation()
			//}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				//code = se.Code
				//reason = se.Reason
			}
			//level, stack := extractError(err)
			//_ = log.WithContext(ctx, logger).Log(level,
				//"kind", "server",
				//"component", kind,
				//"operation", operation,
				//"args", extractArgs(req),
				//"code", code,
				//"reason", reason,
				//"stack", stack,
				//"latency", time.Since(startTime).Seconds(),
			//)
			return
		}
	}
}