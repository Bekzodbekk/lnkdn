package middleware

import (
	"log"
	"net/http"
	"strings"

	t "api-gateway/internal/https/token"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := ctx.Request.URL.Path
		method := ctx.Request.Method

		// Token talab qilinmaydigan endpointlar
		cleanURL := strings.TrimSuffix(url, "/")

		if strings.HasPrefix(cleanURL, "/swagger") ||
			strings.HasPrefix(cleanURL, "/api/v1/auth") ||
			(cleanURL == "/api/v1/users" && method == "POST") {
			ctx.Next()
			return
		}

		// Token tekshirish
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization token is missing Bearer prefix",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := t.ExtractClaim(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Println(claims)

		ctx.Next()
	}
}
