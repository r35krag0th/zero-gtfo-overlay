// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/overlay",
				Handler: ShowAllHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/overlay/expedition",
				Handler: GetExpeditionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/overlay/expedition",
				Handler: UpdateExpeditionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/overlay/sector",
				Handler: GetSectorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/overlay/sector",
				Handler: UpdateSectorHandler(serverCtx),
			},
		},
	)
}
