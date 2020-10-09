package entity

import "github.com/jinzhu/gorm"

type ReadingSchedule struct {
	gorm.Model
	BookId   int
	LastPage int
	DayCount int
}
