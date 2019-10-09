package middleware

import "github.com/gin-gonic/gin"

// Returns slice of all middlewares
func GetAll() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		appMiddleware(),
	}
}
