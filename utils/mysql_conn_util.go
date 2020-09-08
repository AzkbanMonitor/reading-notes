package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

type MysqlConnectionPool struct {
}

var instance *MysqlConnectionPool
var once sync.Once
var db *gorm.DB
var err_db error

func GetInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDBPool() (isSuccess bool) {
	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	dbName := viper.Get("mysql.dbname")
	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	db, err_db := gorm.Open("mysql", username+":"+password+"@("+host+":"+port+")/"+dbName+"?charsetEncoding=utf8&serverTimezone=CTT")
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return true
}
