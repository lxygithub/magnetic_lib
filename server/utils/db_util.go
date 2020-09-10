package utils

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/server/models"
)

func GetXORMEngine() *xorm.Engine {
	//username:password@protocol(address)/dbname?param=value
	username := ReadMySqlConfig("mysql.username")
	password := ReadMySqlConfig("mysql.password")
	url := ReadMySqlConfig("mysql.url")
	dbName := ReadMySqlConfig("mysql.name")

	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8", username, password, url, dbName))

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
	    `enable` TINYINT(1)  NOT NULL DEFAULT 0,
	    `created` DATE NULL DEFAULT NULL,

	    PRIMARY KEY (`id`)
			)


	 */
	engine.Sync2(new(models.Address))

	//engine.Prepare(_sql)
}
