package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func Web() {
	// swager
	facades.Route().Get("/docs/*any", func(ctx http.Context) http.Response {
		handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
		handler(ctx.Response().Writer(), ctx.Request().Origin())
		return nil
	})

	// web
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})
}
