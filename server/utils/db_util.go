package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/server/models"
)

func GetXORMEngine() *xorm.Engine {
	//username:password@protocol(address)/dbname?param=value
	username := ReadMySqlConfig("mysql.username")
	password := ReadMySqlConfig("mysql.password")
	url := ReadMySqlConfig("mysql.url")
	dbName := ReadMySqlConfig("mysql.name")

	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, url, dbName)
	engine, err := xorm.NewEngine("mysql", datasource)

	if err == nil {
		return engine
	}
	return nil
}

func CreateTable() {
	engine := GetXORMEngine()
	defer engine.Close()

	/*
		CREATE TABLE `cili_engine` (
		    `id` INT(10) NOT NULL AUTO_INCREMENT,
		    `url` VARCHAR(64) NULL DEFAULT NULL,
		    `name` VARCHAR(64) NULL DEFAULT NULL,
		    `desc` VARCHAR(64) NULL DEFAULT NULL,
		    `enable` TINYINT(1)  NOT NULL DEFAULT 1,
		    `created` DATE NULL DEFAULT NULL,

		    PRIMARY KEY (`id`)
				)


	*/
	engine.Sync2(new(models.Address))

	//engine.Prepare(_sql)
}
