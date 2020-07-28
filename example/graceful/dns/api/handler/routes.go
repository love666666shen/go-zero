// DO NOT EDIT, generated by goctl
package handler

import (
	"net/http"

	"zero/example/graceful/dns/api/svc"
	"zero/ngin"
)

func RegisterHandlers(engine *ngin.Engine, ctx *svc.ServiceContext) {
	engine.AddRoutes([]ngin.Route{
		{
			Method:  http.MethodGet,
			Path:    "/api/graceful",
			Handler: gracefulHandler(ctx),
		},
	})
}