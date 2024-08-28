package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author string
	Name   string
	Price  string
}
