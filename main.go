package main

import (
	"fmt"
	"net/http"
	"restfullapi/golang/controller"
	"restfullapi/golang/helper"
	"restfullapi/golang/middleware"
	"restfullapi/golang/repository"
	"restfullapi/golang/router"
	"restfullapi/golang/service"

	_ "github.com/go-sql-driver/mysql"

	"restfullapi/golang/config"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := router.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server started on localhost:8080")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
