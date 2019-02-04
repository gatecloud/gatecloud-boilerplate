package middlewares

import "github.com/gin-gonic/gin"

func AddRequestHeader(title, value string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Header.Add(title, value)
		ctx.Next()
	}
}

func AddResponseHeader(title, value string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add(title, value)
		ctx.Next()
	}
}
