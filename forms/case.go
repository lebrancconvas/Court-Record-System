package form

import "gorm.io/gorm"

type Case struct {
	gorm.Model	
	Name string
	Type string
	Location string
	Date string
	Detail string
}