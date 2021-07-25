package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"qiita/model"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@(192.168.99.100:3306)/gin?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init database ok")
}
