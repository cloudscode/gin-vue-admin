package router

import (
	"gin-vue-admin/middleware"
	"gin-vue-admin/modules/research/paper/api"
	"github.com/gin-gonic/gin"
)

func InitResearchRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("paper").Use(middleware.OperationRecord())
	{
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表

	}
}
