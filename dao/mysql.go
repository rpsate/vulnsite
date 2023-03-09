package dao

import (
	"vulnsite/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDb() {
	driver := utils.DbDrive
	host := utils.DbHost
	port := utils.DbPort
	username := utils.DbUsername
	password := utils.DbPassword
	database := utils.DbDatabase
	charset := utils.DbCharset

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	var err error
	DB, err = gorm.Open(driver, args)
	if err != nil {
		panic(err.Error())
	}
	err = DB.DB().Ping()
	if err != nil {
		panic(err.Error())
	}
}

func Close() {
	DB.Close()
}