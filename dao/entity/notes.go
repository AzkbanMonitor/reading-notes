package entity

import "github.com/jinzhu/gorm"

type Notes struct {
	gorm.Model
	BookId  int
	Details string
}
