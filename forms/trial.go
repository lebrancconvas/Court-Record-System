package form

import "gorm.io/gorm"

type Trial struct {
	gorm.Model
	Detail string
	Date string
}