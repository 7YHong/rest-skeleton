package utils

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"time"
)

var Logger *zap.Logger

const (
	LOG_NAME = "flyPenguin"
	LOG_DIR  = "logs"
)

func init() {
	if _, err := os.Stat(LOG_DIR); err != nil {
		if err := os.Mkdir(LOG_DIR, os.ModePerm); err != nil {
			log.Fatalln("无法创建日志目录")
		}
	}
	core := zapcore.NewCore(getEncoder(),
		//zapcore.AddSync(os.Stdout),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter())),
		zapcore.DebugLevel)

	Logger = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriter() io.Writer {
	filename := "logs/" + LOG_NAME
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(24*30*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	//hook, err := rotatelogs.New(
	//	filename+"_%Y%m%d%S.log",
	//	rotatelogs.WithLinkName(filename),
	//	rotatelogs.WithMaxAge(60*time.Second),
	//	rotatelogs.WithRotationTime(10*time.Second),
	//)
	if err != nil {
		panic(err)
	}
	return hook
}
