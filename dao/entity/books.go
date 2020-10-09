package entity

import (
	"github.com/jinzhu/gorm"
)

type Books struct {
	gorm.Model
	BookName   string
	Writer     string
	TotalPages int
	Status     bool
}
