// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"Hermes/api/inter/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hermesen",
				Handler: HermesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: ExpandHandler(serverCtx),
			},
		},
	)
}
