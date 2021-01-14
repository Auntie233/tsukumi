package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
	"tsukumi/app/config"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.LogFilePath
	logFileName := config.LogFileName

	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, err := rotateLogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotateLogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(lfHook)
	return func(context *gin.Context) {
		startTime := time.Now()
		// 处理请求
		context.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := context.Request.Method
		reqUri := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		logger.Infof("|%3d|%13v|%15s|%s|%s|",
			statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
}
