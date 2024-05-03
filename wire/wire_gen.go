// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/corysakti/cats-social-go/app/modules/hellowordls"
	"github.com/corysakti/cats-social-go/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// Injectors from wire.go:

func SetupApp(db *sqlx.DB) *gin.Engine {
	helloWorldRepositories := hellowordls.NewHelloWorldRepositories(db)
	helloWorldServices := hellowordls.NewHelloWorldServices(helloWorldRepositories)
	helloWorldHandlers := hellowordls.NewHelloWorldHandlers(helloWorldServices)
	engine := routes.Routes(helloWorldHandlers)
	return engine
}

// wire.go:

var setVariable = wire.NewSet(hellowordls.NewHelloWorldRepositories, hellowordls.NewHelloWorldServices, hellowordls.NewHelloWorldHandlers)
