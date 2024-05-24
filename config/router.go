package config

import (
	"fmt"
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(handlers ...interface{}) *httprouter.Router {
	router := httprouter.New()

	for _, handler := range handlers {
		switch h := handler.(type) {
		case controller.CategoryController:
			router.GET("/api/categories", h.FindAll)
			router.GET("/api/categories/:categoryId", h.FindById)
			router.POST("/api/categories", h.Create)
			router.PUT("/api/categories/:categoryId", h.Update)
			router.DELETE("/api/categories/:categoryId", h.Delete)
		case controller.AuthenticationController:
			router.GET("/v1/user/register", h.Register)
			router.GET("/v1/user/login", h.Login)
		default:
			// Handle unsupported handler types
			panic(fmt.Sprintf("Unsupported handler type: %T", h))
		}
	}

	router.PanicHandler = exception.ErrorHandler

	return router
}
