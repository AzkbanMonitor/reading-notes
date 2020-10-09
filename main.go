package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "reading-notes/config"
	"reading-notes/utils"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)

}
func main() {
	//log.SetOutput(os.Stdout)
	//config.InitViper()
	get := viper.Get("host")
	log.Info(get)
	isSuccess := utils.GetInstance().InitDBPool()
	if !isSuccess {
		log.Error("db init err")
		//os.Exit(1)
	}
	router := gin.Default()
	router.Run(":8080")
}
