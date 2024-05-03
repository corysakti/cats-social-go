package main

import (
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/database"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/repository/impl"
	impl2 "github.com/corysakti/cats-social-go/service/impl"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq" // PostgreSQL driver
	"net/http"
)

func main() {

	db := database.NewDB()
	validate := validator.New()
	categoryRepository := impl.NewCategoryRepositoryImpl()
	categoryService := impl2.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
