package system

import (
	v1 "github.com/Sectran/sectran_admin/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("listUser", userApi.ListUser)
		userRouter.POST("addUser", userApi.AddUser)
		userRouter.POST("updateUser", userApi.UpdateUser)
		userRouter.POST("deleteUser", userApi.DeleteUser)
	}
}
