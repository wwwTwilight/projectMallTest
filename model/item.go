package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Owner       string `json:"owner"`
}
