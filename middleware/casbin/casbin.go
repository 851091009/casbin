package casbin

import (
	expansionCasbin "feita/expansion/casbin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinAuth(notCheckLoginUrlArr []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 验证权限
		result, err := expansionCasbin.Enforcer.Enforce("1", "/v1/admin/index/index", "GET")

		if err != nil {
			// 验证错误
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "权限验证错误，代码：002，错误信息：" + err.Error(),
			})
			ctx.Abort()
			return
		}
		// 判断权限
		if result != true {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "没有访问权限",
			})
			ctx.Abort()
			return
		}

		// 鉴权完毕
		ctx.Next()
	}
}
