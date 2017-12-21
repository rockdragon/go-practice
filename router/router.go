package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router 路由
func Router(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.Run(":8088")
}
