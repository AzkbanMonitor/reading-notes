package middleware

import (
	"github.com/gin-gonic/gin"
	"reading-notes/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if error := recover(); error != nil {
				error := errors.SysError
				context.JSON(200, error)
				return
			}
		}()
		context.Next()
	}
}
