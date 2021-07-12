/*
* Initialize mysql database
* Author: Lv Wenchao
* Date: 2021-07-12
* Last Modified:
* Last Modified Date:
 */
package models

import (
	"fmt"
	"log"

	"cssd-admin/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to load section 'database' %v", err)
	}

	// dbType := sec.Key("TYPE").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	dbName := sec.Key("NAME").String()
	// tablePrefix := sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Fail to open mysql server :\n%v", err)
	}
}
