package main

import (
	"fmt"
	"net/http"
	"omclabs/go-qontak/app/handlers/controllers/qontak_controller"
	"omclabs/go-qontak/app/handlers/middlewares"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/models/web"
	"omclabs/go-qontak/app/repositories/qontak_repository"
	"omclabs/go-qontak/app/services/qontak_service"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

var logError web.LogError
var logRequest web.LogRequest
var logResponse web.LogResponse

func main() {
	err := godotenv.Load()
	if err != nil {
		logError.ErrorDesc = err.Error()
		helpers.WriteLog("fatal", "error", "failed to load env file, terminating process", logRequest, logResponse, logError)
	}

	// validate := validator.New()
	// db := config.NewMysqlConn()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	crmRepository := qontak_repository.NewCrmRepository()
	crmService := qontak_service.NewCrmService(crmRepository, client)
	crmController := qontak_controller.NewCrmController(crmService)

	chatRepository := qontak_repository.NewChatRepository()
	chatService := qontak_service.NewChatService(chatRepository, client)
	chatController := qontak_controller.NewChatController(chatService)

	router := httprouter.New()

	router.GET("/api/v1/crm/get-params", crmController.GetParam)
	router.GET("/api/v1/crm/contacts", crmController.GetContact)
	router.POST("/api/v1/crm/contacts", crmController.CreateContact)
	router.GET("/api/v1/crm/contacts/:contact_id", crmController.GetContactById)
	router.PUT("/api/v1/crm/contacts/:contact_id", crmController.UpdateContact)
	router.DELETE("/api/v1/crm/contacts/:contact_id", crmController.DeleteContact)

	router.GET("/api/v1/omnichannel/wa-integrations", chatController.GetWhatsappIntegration)
	router.GET("/api/v1/omnichannel/contact-list", chatController.GetContactList)
	router.GET("/api/v1/omnichannel/wa-templates", chatController.GetWhatsappTemplates)
	router.POST("/api/v1/omnichannel/validate-number", chatController.ValidateNumber)

	// router.GET("/api/omnichannel/integrations", chatController.GetIntegrations)
	// router.GET("/api/omnichannel/integrations/:channel", chatController.GetIntegrationsByChannel)
	// router.GET("/api/omnichannel/contact-list", chatController.GetContactList)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(helpers.NewNotFoundError())
	})
	router.PanicHandler = helpers.ErrorHandler
	server := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: middlewares.NewGeneralMiddleware(router),
		// Handler: router,
	}
	helpers.WriteLog("info", "event", fmt.Sprintf(`starting server at port : %s`, os.Getenv("APP_PORT")), logRequest, logResponse, logError)

	err = server.ListenAndServe()
	if err != nil {
		logError.ErrorDesc = err.Error()
		helpers.WriteLog("fatal", "error", "failed to start server, terminating process", logRequest, logResponse, logError)
	}
}
