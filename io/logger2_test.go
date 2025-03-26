package io_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOG      = "这里是日志内容"
	TIME_FMT = "2006-01-02 15:04:05.000"
)

var (
	loc *time.Location
)

func init() {
	loc, _ = time.LoadLocation("asia/shanghai")
}

func BenchmarkZap(b *testing.B) {
	logFile := "../log/zap.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(TIME_FMT) //指定时间格式
	encoderConfig.TimeKey = "time"                                   //默认是ts
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder          //指定level的显示样式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), //json格式
		zapcore.AddSync(fout),                 //指定输出到文件
		zapcore.InfoLevel,                     //设置最低级别
	)
	logger := zap.New(
		core,
		zap.AddCaller(), //上报文件名和行号
	)
	logger = logger.With(
		zap.Namespace("uber"), //后续的Field都记录在此命名空间中
		//通过zap.String、zap.Int等显式指定类型；fmt.Printf之类的方法大量使用interface{}和反射，性能损失不少
		zap.String("bizID", "123456"), //公共的Field
	)

	b.ResetTimer() //基准测试开始计时
	for b.Loop() {
		logger.Info(LOG, zap.String("name", "dqq"), zap.Int("age", 18))
	}
	logger.Sync() //把缓冲里的内容刷入磁盘
}

func BenchmarkSlog(b *testing.B) {
	logFile := "../log/slog.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	handler := slog.NewJSONHandler( //json格式
		fout, //指定输出到文件
		&slog.HandlerOptions{
			AddSource: true,           //上报文件名和行号
			Level:     slog.LevelInfo, //设置最低级别
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey { //如果Key=="time"
					t := a.Value.Time()
					a.Value = slog.StringValue(t.Format(TIME_FMT)) //替换Value
				}
				return a
			},
		},
	)
	logger := slog.New(handler)

	logger = logger.
		WithGroup("google").                 //分组
		With(slog.String("bizId", "123456")) //公共的Field

	b.ResetTimer() //基准测试开始计时
	for b.Loop() {
		logger.Info(LOG, slog.String("name", "dqq"), slog.Int("age", 18))
	}
}

func BenchmarkLogrus(b *testing.B) {
	logFile := "../log/logrus.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetOutput(fout)                     //设置日志文件
	logger.SetLevel(logrus.InfoLevel)          //设置最低级别
	logger.SetReportCaller(true)               //输出是从哪里调起的日志打印，日志里会包含func和file
	logger.SetFormatter(&logrus.JSONFormatter{ //JSON格式
		TimestampFormat: TIME_FMT, //时间格式
	})

	logEntry := logger.WithField("bizId", "123456") //公共的Field
	//logrus没有分组的功能

	b.ResetTimer() //基准测试开始计时
	for b.Loop() {
		logEntry.WithFields(logrus.Fields{"name": "dqq", "age": 18}).Info(LOG)
	}
}

// go test ./io -bench=^BenchmarkZap$ -run=^$
// go test ./io -bench=^BenchmarkSlog$ -run=^$
// go test ./io -bench=^BenchmarkLogrus$ -run=^$

/**
goos: windows
goarch: amd64
pkg: dqq/go/frame/io
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz

普通文本格式
BenchmarkZap-8            324475              3756 ns/op
BenchmarkSlog-8           297176              3996 ns/op
BenchmarkLogrus-8         141416              8108 ns/op

json格式
BenchmarkZap-8            301575              4318 ns/op
BenchmarkSlog-8           267660              4556 ns/op
BenchmarkLogrus-8         130428              8356 ns/op
*/
