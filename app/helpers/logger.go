package helpers

import (
	"io"
	"omclabs/go-qontak/app/models/web"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	isPrettyPrint := true
	if os.Getenv("APP_ENV") == "production" {
		isPrettyPrint = false
	}

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: isPrettyPrint,
	})
}

func Logger(level string, filename string, message string, loggerData web.LoggerData) {
	dateNow := time.Now().Format("20060102") //YmD
	path := "./logs/" + dateNow + "/"
	fileName := filename + ".json"
	isExists := CheckDirOrFileExists(path)
	if !isExists {
		os.MkdirAll(path, 0777)
	}

	filePath := path + fileName
	files, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer files.Close()

	objFields := StructToMap(loggerData)
	writer := io.MultiWriter(os.Stdout, files)
	log.SetOutput(writer)

	switch level {
	case "warn":
		log.WithFields(objFields).Warn(message)
	case "error":
		log.WithFields(objFields).Error(message)
	case "panic":
		log.WithFields(objFields).Panic(message)
	case "fatal":
		log.WithFields(objFields).Fatal(message)
	default:
		log.WithFields(objFields).Info(message)
	}
}

func WriteLog(level string, filename string, message string, LogRequest web.LogRequest, LogResponse web.LogResponse, logError web.LogError) {
	loggerData := web.LoggerData{
		Error:    logError,
		Request:  LogRequest,
		Response: LogResponse,
	}

	Logger(level, filename, message, loggerData)
}
