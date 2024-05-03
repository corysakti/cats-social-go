package hellowordls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorldHandlers struct {
	serv *HelloWorldServices
}

func NewHelloWorldHandlers(serv *HelloWorldServices) *HelloWorldHandlers {
	return &HelloWorldHandlers{serv: serv}
}

func (h *HelloWorldHandlers) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Worlds",
	})
}
