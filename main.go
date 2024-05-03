package main

import (
	"github.com/corysakti/cats-social-go/config"
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/database"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/middleware"
	"github.com/corysakti/cats-social-go/repository/impl"
	impl2 "github.com/corysakti/cats-social-go/service/impl"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq" // PostgreSQL driver
	"net/http"
)

func main() {

	db := database.NewDB()
	validate := validator.New()
	categoryRepository := impl.NewCategoryRepositoryImpl()
	categoryService := impl2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := config.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
