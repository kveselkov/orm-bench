package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID          int
	Name        string
	Phone       string
	Date        string
	City        string
	Country     string
	Email       string
	Coordinates string
}

type UserDB struct {
	DB *gorm.DB
}

func (u *UserDB) Fetch(limit int) ([]*User, error) {
	var result []*User
	err := u.DB.Limit(limit).Find(&result).Error

	return result, err
}
