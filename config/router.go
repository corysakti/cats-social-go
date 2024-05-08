package config

import (
	"fmt"
	"github.com/corysakti/cats-social-go/controller"
	"github.com/corysakti/cats-social-go/exception"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(handlers ...interface{}) *httprouter.Router {
	router := httprouter.New()

	for _, handler := range handlers {
		switch h := handler.(type) {
		case controller.CategoryController:
			// Add routes for the category controller
			router.GET("/api/categories", h.FindAll)
			router.GET("/api/categories/:categoryId", h.FindById)
			router.POST("/api/categories", h.Create)
			router.PUT("/api/categories/:categoryId", h.Update)
			router.DELETE("/api/categories/:categoryId", h.Delete)
		case func(http.ResponseWriter, *http.Request):
			// Add custom routes for each additional handler
			router.HandlerFunc("GET", "/api/customroute1", h)
			// Add more custom routes as needed
		// Add more cases for different types of handlers
		default:
			// Handle unsupported handler types
			panic(fmt.Sprintf("Unsupported handler type: %T", h))
		}
	}

	router.PanicHandler = exception.ErrorHandler

	return router
}
