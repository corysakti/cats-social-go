//go:build wireinject

// + build:wireinject
package wire

import (
	"github.com/corysakti/cats-social-go/app/modules/hellowordls"
	"github.com/corysakti/cats-social-go/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var setVariable = wire.NewSet(
	// Repositories
	hellowordls.NewHelloWorldRepositories,

	// Services
	hellowordls.NewHelloWorldServices,

	// Handlers
	hellowordls.NewHelloWorldHandlers,
)

func SetupApp(db *sqlx.DB) *gin.Engine {
	wire.Build(
		setVariable,
		routes.Routes,
	)

	return nil
}
