package main

import (
	"log"
	"net/http"
	config "omclabs/go-qontak/app/configs/database"
	"omclabs/go-qontak/app/handlers/controllers"
	"omclabs/go-qontak/app/handlers/controllers/qontak_controller"
	"omclabs/go-qontak/app/handlers/exceptions"
	"omclabs/go-qontak/app/helpers"
	"omclabs/go-qontak/app/repositories"
	"omclabs/go-qontak/app/repositories/qontak_repository"
	"omclabs/go-qontak/app/services"
	"omclabs/go-qontak/app/services/qontak_service"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Cant load env file, terminating")
	}

	validate := validator.New()
	db := config.NewMysqlConn()

	var client *http.Client
	client = &http.Client{
		Timeout: 30 * time.Second,
	}

	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	crmRepository := qontak_repository.NewCrmRepository()
	crmService := qontak_service.NewCrmService(crmRepository, client)
	crmController := qontak_controller.NewCrmController(crmService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/crm/get-params", crmController.GetParam)
	router.GET("/api/crm/contacts", crmController.GetContact)
	router.GET("/api/crm/contacts/:id", crmController.GetContactById)
	router.POST("/api/crm/contacts", crmController.CreateContact)
	router.DELETE("/api/crm/contacts/:id", crmController.DeleteContact)
	router.PUT("/api/crm/contacts/:id", crmController.UpdateContact)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(exceptions.NewNotFoundError("Page not found"))
	})
	router.PanicHandler = exceptions.ErrorHandler

	server := http.Server{
		Addr: "localhost:" + os.Getenv("APP_PORT"),
		// Handler: middleware.NewAuthMiddleware(router),
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
