package account

import "github.com/gin-gonic/gin"

func Urlpatterns(router *gin.RouterGroup) {
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": "ok"})
	})
	router.GET("/userInfo")
}
