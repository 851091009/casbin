package routers

import (
	"feita/controllers"
	"feita/helpers"
	"feita/middleware/casbin"
	expansionCasbin "feita/expansion/casbin"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/v1/admin", controllers.Index)
	AuthGroup := router.Group(helpers.ApiAdminPrefix)

	// 免验证URL
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, helpers.ApiAdminPrefix+"/roleMember/memberRoleList")

	// casbin 初始化
	expansionCasbin.CasbinInit()

	// token 验证、权限管理
	AuthGroup.Use(casbin.CasbinAuth(notCheckLoginUrlArr))
	{
		// 部门管理
		AuthGroup.GET("/index/index", controllers.Index)         // 部门列表
	}
	return router
}
