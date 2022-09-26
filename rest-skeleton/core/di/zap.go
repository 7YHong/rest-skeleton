package di

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xdi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	_ "rest-skeleton/core/dotenv"
	"time"
)

var LOG_NAME string

func init() {
	LOG_NAME = dotenv.Getenv("APP_NAME").String("rest-skeleton")
	obj := xdi.Object{
		Name: "zap",
		New: func() (i interface{}, e error) {
			core := zapcore.NewCore(
				zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
					TimeKey:       "T",
					LevelKey:      "L",
					NameKey:       "N",
					CallerKey:     "C",
					MessageKey:    "M",
					StacktraceKey: "S",
					LineEnding:    zapcore.DefaultLineEnding,
					EncodeLevel:   zapcore.CapitalLevelEncoder,
					EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
						enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
					},
					EncodeDuration: zapcore.StringDurationEncoder,
					EncodeCaller:   zapcore.ShortCallerEncoder,
				}),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter())),
				zap.InfoLevel,
			)
			logger := zap.New(core, zap.AddCaller())
			return logger.Sugar(), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Zap() (logger *zap.SugaredLogger) {
	if err := xdi.Populate("zap", &logger); err != nil {
		panic(err)
	}
	return
}

type ZapOutput struct {
	Logger *zap.SugaredLogger
}

func (t *ZapOutput) Write(p []byte) (n int, err error) {
	t.Logger.Info(string(p))
	return len(p), nil
}

func getWriter() io.Writer {
	filename := "logs/" + LOG_NAME
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
