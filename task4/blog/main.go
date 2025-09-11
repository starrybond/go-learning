package main

import (
	"blog/model"
	"blog/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	model.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	utils.InitLogger()
	model.Migrate(model.DB)
}

func main() {
	r := NewRouter()
	utils.L.Info("start on :8080")
	r.Run(":8080")
}
