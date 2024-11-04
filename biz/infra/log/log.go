package log

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

func InitLog() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic("创建日志文件夹失败：" + err.Error())
	}

	logFileName := filepath.Join("logs", "log-"+time.Now().Format("2006-01-02-15-04-05")+".log")
	// 打开一个文件用于写日志
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		klog.Fatalf("failed to open log file: %v", err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	klog.SetOutput(multiWriter)
}
