package models

import "gorm.io/gorm"

type Shortcut struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"` //name e.g Ctrl+x
	For         string `json:"for"`                //which IDE
	Description string `json:"usage"`
}
