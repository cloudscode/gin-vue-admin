package utils

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"github.com/gin-gonic/gin"
)

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID.String()
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserCustomClaims(c *gin.Context) CustomClaims {
//	if claims, exists := c.Get("claims"); !exists {
//		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
//		return
//	} else {
//		waitUse := claims.(*request.CustomClaims)
//		return waitUse
//	}
//}
