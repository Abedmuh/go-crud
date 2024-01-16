package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"belajar-go/go-crud/app"
	"belajar-go/go-crud/controller"
	"belajar-go/go-crud/helper"
	"belajar-go/go-crud/middleware"
	"belajar-go/go-crud/repository"
	"belajar-go/go-crud/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	collectionRepository := repository.NewCollectionRepository()
	collectionService := service.NewCollectionService(collectionRepository, db, validate)
	collectionController := controller.NewCollectionController(collectionService)
	router := app.NewRouter(collectionController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
