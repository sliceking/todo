package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Description string
	Done        bool
}
