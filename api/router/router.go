package router

import (
	"github.com/gin-gonic/gin"
	"reading-notes/api/controller"
	"reading-notes/api/middleware"
)

func ApiRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	apiRouter := r.Group("/api/v1")
	{
		apiRouter.POST("/addBooks", controller.AddBooks)
	}

}
