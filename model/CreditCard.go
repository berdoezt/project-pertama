package model

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	CardNumber string
	PersonUUID string
}
