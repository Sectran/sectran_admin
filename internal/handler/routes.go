// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	access_policy "sectran_admin/internal/handler/access_policy"
	account "sectran_admin/internal/handler/account"
	base "sectran_admin/internal/handler/base"
	department "sectran_admin/internal/handler/department"
	device "sectran_admin/internal/handler/device"
	role "sectran_admin/internal/handler/role"
	user "sectran_admin/internal/handler/user"
	"sectran_admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/department/create",
					Handler: department.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/update",
					Handler: department.UpdateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/delete",
					Handler: department.DeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/list",
					Handler: department.GetDepartmentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department",
					Handler: department.GetDepartmentByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/delete",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.GetRoleByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.GetUserByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/account/create",
					Handler: account.CreateAccountHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/account/update",
					Handler: account.UpdateAccountHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/account/delete",
					Handler: account.DeleteAccountHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/account/list",
					Handler: account.GetAccountListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/account",
					Handler: account.GetAccountByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/device/create",
					Handler: device.CreateDeviceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device/update",
					Handler: device.UpdateDeviceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device/delete",
					Handler: device.DeleteDeviceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device/list",
					Handler: device.GetDeviceListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device",
					Handler: device.GetDeviceByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)


	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/access_policy/create",
					Handler: access_policy.CreateAccessPolicyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/access_policy/update",
					Handler: access_policy.UpdateAccessPolicyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/access_policy/delete",
					Handler: access_policy.DeleteAccessPolicyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/access_policy/list",
					Handler: access_policy.GetAccessPolicyListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/access_policy",
					Handler: access_policy.GetAccessPolicyByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
