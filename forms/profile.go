package form

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name string
	ImageURL string
	Age int
	Gender string
	Detail string
}