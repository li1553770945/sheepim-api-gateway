package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

func InitLog() *logrus.Logger {
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic("创建日志文件夹失败：" + err.Error())
	}

	logFileName := filepath.Join("logs", "log-"+time.Now().Format("2006-01-02-15-04-05")+".log")
	// 打开一个文件用于写日志
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		hlog.Fatalf("无法打开日志文件: %v", err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	hlog.SetOutput(multiWriter)
	return logger
}
