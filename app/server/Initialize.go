package server

import (
	"fmt"
	"log"
	"tpk-backend/app/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (h *FuncHandler) Initialize() {
	db := config.LoadDB()
	dns := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local`, db.Username, db.Password, db.Host, db.Port, db.Database)
	log.Println(dns)
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = conn
}