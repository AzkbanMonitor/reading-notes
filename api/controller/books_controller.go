package controller

import (
	"github.com/gin-gonic/gin"
	"reading-notes/dto"
)

func AddBooks(ctx *gin.Context) {
	var booksDto dto.Books
	if err := ctx.ShouldBindJSON(&booksDto); err != nil {
		ctx.JSON(400, "参数异常")
		return
	}

}
