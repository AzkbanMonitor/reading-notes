package main

import (
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	router := gin.Default()
	router.Run(":8080")
}
