package form

import "gorm.io/gorm"

type Evidence struct {
	gorm.Model
	Name string	
	ImageURL string
	Type string
	Description string
	Source string
}
