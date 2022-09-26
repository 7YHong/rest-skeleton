package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rest-skeleton/utils"
	"runtime"
	"strconv"
)

func Recovery(c *gin.Context, rec interface{}) {
	logger := utils.Logger
	c.String(500, "%v", rec)
	c.Abort()
	logger.Sugar().Warn(rec)
	printStack(logger)
}

func printStack(logger *zap.Logger) {
	pcs := make([]uintptr, 10)
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	var byt bytes.Buffer
	byt.WriteString("\n")
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		byt.WriteString(frame.Function)
		byt.WriteString("\n\t")
		byt.WriteString(frame.File)
		byt.WriteString(":")
		byt.WriteString(strconv.Itoa(frame.Line))
		byt.WriteString("\n")
	}
	logger.Warn(byt.String())
}
