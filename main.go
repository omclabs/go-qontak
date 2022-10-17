package main

import (
	"net/http"
	"omclabs/go-qontak/app/handlers/controllers/qontak_controller"
	"omclabs/go-qontak/app/handlers/middlewares"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/repositories/qontak_repository"
	"omclabs/go-qontak/app/services/qontak_service"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	errEnv := godotenv.Load()
	helpers.PanicIfError(errEnv)

	// validate := validator.New()
	// db := config.NewMysqlConn()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	crmRepository := qontak_repository.NewCrmRepository()
	crmService := qontak_service.NewCrmService(crmRepository, client)
	crmController := qontak_controller.NewCrmController(crmService)

	router := httprouter.New()

	router.GET("/api/crm/get-params", crmController.GetParam)
	router.GET("/api/crm/contacts", crmController.GetContact)
	router.GET("/api/crm/contacts/:id", crmController.GetContactById)
	router.POST("/api/crm/contacts", crmController.CreateContact)
	router.DELETE("/api/crm/contacts/:id", crmController.DeleteContact)
	router.PUT("/api/crm/contacts/:id", crmController.UpdateContact)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(123)
	})
	router.PanicHandler = helpers.ErrorHandler

	server := http.Server{
		Addr:    "localhost:" + os.Getenv("APP_PORT"),
		Handler: middlewares.NewGeneralMiddleware(router),
		// Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
