package handler

import (
	"net/http"
	base "sectran_admin/internal/handler/base"
	department "sectran_admin/internal/handler/department"
	"sectran_admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlersCustom(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: base.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update_authority_api",
				Handler: base.UpdateApiAuthorityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/department/children",
				Handler: department.GetChDepartmentListHandler(serverCtx),
			},
		},
	)
}
