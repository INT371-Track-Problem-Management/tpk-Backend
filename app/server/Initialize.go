package server

import (
	"context"
	"fmt"
	"log"
	"tpk-backend/app/pkg/config"

	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (h *FuncHandler) InitContext() {
	h.ctx = context.Background()
}

func (h *FuncHandler) InitDatabase() {
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

func (h *FuncHandler) InitFile() {

	sa := option.WithCredentialsFile("../../creadential/rungmod-senior-project-firebase-adminsdk-yhyb5-d43c34d5c2.json")
	app, err := firebase.NewApp(h.ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	h.client, err = app.Firestore(h.ctx)
	if err != nil {
		log.Fatalln(err)
	}

	h.storage, err = cloud.NewClient(h.ctx, sa)
	if err != nil {
		log.Fatalln(err)
	}
}
