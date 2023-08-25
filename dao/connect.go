package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	// 设置 Viper 的配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config:", err)
	}

	// 连接到数据库
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.hostname")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, hostname, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
