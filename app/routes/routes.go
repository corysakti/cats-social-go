package routes

import (
	"github.com/corysakti/cats-social-go/app/modules/hellowordls"
	"github.com/gin-gonic/gin"
)

func Routes(helloWorld *hellowordls.HelloWorldHandlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.GET("/hello", helloWorld.HelloWorld)

	return r
}
