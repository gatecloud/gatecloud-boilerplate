package middlewares

import (
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors sets cross origin parameters
func Cors(enabled bool) gin.HandlerFunc {
	if enabled {
		return cors.New(cors.Config{
			AllowMethods:     []string{"DELETE", "PATCH", "POST", "GET", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return regexp.MustCompile("(wisyedu.com)|(timingniao.com)|(ioi.nz)").
					MatchString(origin)
			},
			MaxAge: 12 * time.Hour,
		})
	}
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
