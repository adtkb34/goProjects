package main

import (
	"awesomeProject/account"
	"awesomeProject/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库
	db, _ := dao.ConnectDB()
	//数据库迁移
	dao.MigrateDB(db)
	router := gin.Default()
	account.Urlpatterns(router.Group("account"))
	_ = router.Run()
}
