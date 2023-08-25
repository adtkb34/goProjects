package dao

import (
	"awesomeProject/account"
	"awesomeProject/app1"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

func MigrateDB(db *gorm.DB) {
	//运行数据库迁移
	modelArrs := []interface{}{
		account.AllModels(),
		app1.AllModels(),
	}
	//失败，判断是否有重复的表名，modelArrs=[[0xc00068d050 0xc000009038] [0xc00068d080 0xc000009050]]
	modelMap := make(map[interface{}]bool)
	for _, modelArr := range modelArrs {
		modelArr := modelArr
		go func() {
			for _, model := range modelArr.([]interface{}) {
				fmt.Println(reflect.TypeOf(model))
				if _, ifExist := modelMap[model]; !ifExist {
					modelMap[model] = true
				} else {
					panic("有模型重复")
				}
			}
		}()
	}

	for _, modelArr := range modelArrs {
		modelArr := modelArr
		go func() {
			err := db.AutoMigrate(modelArr.([]interface{})...)
			if err != nil {
				return
			}
		}()
	}

	//done := make(chan int) // 通过通道进行协程同步
	//modelArrCh := make(chan []interface{})
	//go func() {
	//	modelArrCh <- account.AllModels()
	//
	//	done <- 1
	//}()
	//
	//modelMap := make(map[interface{}]bool)
	//for modelArr := range modelArrCh {
	//	modelArr := modelArr
	//	go func() {
	//		for _, model := range modelArr {
	//			if _, ifExist := modelMap[model]; !ifExist {
	//				modelMap[model] = true
	//			} else {
	//				panic("有模型重复")
	//			}
	//		}
	//	}()
	//}
	//<-done
	//go func() {
	//	for modelArr := range modelArrCh {
	//		modelArr := modelArr
	//		go func() {
	//			err := db.AutoMigrate(modelArr...)
	//			if err != nil {
	//				return
	//			}
	//		}()
	//	}
	//}()
}
