package handler

import (
	"net/http"
	base "sectran_admin/internal/handler/base"
	"sectran_admin/internal/handler/file"
	"sectran_admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlersCustom(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/update_authority_api",
					Handler: base.UpdateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/get_menu_authority_list",
					Handler: base.GetMenuAuthorityHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: base.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: file.UploadHandler(serverCtx),
				},
			}...,
		),
	)
}
