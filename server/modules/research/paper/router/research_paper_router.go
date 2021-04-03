package router

import (
	"gin-vue-admin/middleware"
	"gin-vue-admin/modules/research/paper/api"
	"github.com/gin-gonic/gin"
)

func InitResearchRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("paper").Use(middleware.OperationRecord())
	{
		UserRouter.POST("getUserList", v1.GetUserList)           // 列表
		UserRouter.POST("createInfo", v1.Create)           // 新增
		UserRouter.PUT("updateInfo", v1.Update)           // 编辑
		UserRouter.GET("getInfo", v1.Get)           // 获取一条
		UserRouter.DELETE("deleteInfo", v1.Delete)           // 删除
	}
}
