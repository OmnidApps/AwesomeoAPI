package routes

import (
	"apiTemplate/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.PingAPI{}.Get)
}
