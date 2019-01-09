package xorm

import (
	"github.com/go-xorm/xorm"
)

type Users struct {
	Id          int64
	Name        string
	Phone       string
	Date        string
	City        string
	Country     string
	Email       string
	Coordinates string
}

type UserDB struct {
	DB *xorm.Engine
}

func (u *UserDB) Fetch(limit int) ([]Users, error) {
	users := make([]Users, 0)
	err := u.DB.Limit(limit).Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
