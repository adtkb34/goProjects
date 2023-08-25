package app1

import "github.com/gin-gonic/gin"

func Urlpatterns(router *gin.RouterGroup) {
	router.GET("/example")
}
