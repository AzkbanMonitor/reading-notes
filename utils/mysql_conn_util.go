package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type MysqlConnectionPool struct {
}

var instance *MysqlConnectionPool
var once sync.Once
var db *gorm.DB
var errDb error

func GetInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDBPool() (isSuccess bool) {
	host := viper.Get("host")
	port := viper.Get("port")
	dbName := viper.Get("dbname")
	username := viper.Get("username")
	password := viper.Get("password")
	connUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", username, password, host, port, dbName)
	log.Info(connUrl)
	db, errDb := gorm.Open("mysql", connUrl)
	if errDb != nil {
		log.Error(errDb)
		return false
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return true
}

func GetDb() *gorm.DB {
	return db
}
