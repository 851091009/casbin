package middleware

import "github.com/gin-gonic/gin"

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// 设置定义中间件跳过函数
func AllowPath(path string, prefixes []string) bool {
	for _, p := range prefixes {
		if p == path {
			return true
		}
	}
	return false
}
