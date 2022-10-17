package helpers

import (
	"io"
	"net/http"
	"omclabs/go-qontak/app/models/web"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

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
	writer := io.MultiWriter(os.Stdout, files)

	logger := logrus.New()
	isPrettyPrint := false
	if os.Getenv("APP_ENV") == "development" {
		isPrettyPrint = true
	}

	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: isPrettyPrint,
	})

	logger.SetOutput(writer)

	objFields := StructToMap(loggerData)

	switch level {
	case "warn":
		logger.WithFields(objFields).Warn(message)
	case "error":
		logger.WithFields(objFields).Error(message)
	case "panic":
		logger.WithFields(objFields).Panic(message)
	default:
		logger.WithFields(objFields).Info(message)
	}
}

func EventLogger(writer http.ResponseWriter, request *http.Request, payload interface{}, body web.ApiResponse) {
	loggerData := web.LoggerData{
		Request: struct {
			Url     string      "json:\"url\""
			Header  http.Header "json:\"header\""
			Method  string      "json:\"method,omitempty\""
			Payload interface{} "json:\"payload,omitempty\""
		}{
			Url:     request.RequestURI,
			Header:  request.Header,
			Method:  request.Method,
			Payload: payload,
		},
		Response: struct {
			Code   int         "json:\"code,omitempty\""
			Status string      "json:\"status,omitempty\""
			Header http.Header "json:\"header,omitempty\""
			Body   interface{} "json:\"body,omitempty\""
			Error  string      "json:\"error,omitempty\""
		}{
			Code:   body.Code,
			Status: body.Status,
			Header: writer.Header(),
			Body:   body,
		},
	}

	Logger("info", "event", "serving request", loggerData)
}
