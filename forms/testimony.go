package form

import "gorm.io/gorm"

type Testimony struct {
	gorm.Model
	Declarator string
	Statement []string
}
