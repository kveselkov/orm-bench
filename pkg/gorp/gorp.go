package gorp

import (
	"fmt"

	"gopkg.in/gorp.v1"
)

type Users struct {
	ID          int32  `db:"id"`
	Name        string `db:"name"`
	Phone       string `db:"phone"`
	Date        string `db:"date"`
	City        string `db:"city"`
	Country     string `db:"country"`
	Email       string `db:"email"`
	Coordinates string `db:"coordinates"`
}

type UserDB struct {
	DB *gorp.DbMap
}

func (u *UserDB) Fetch(limit int) ([]Users, error) {
	users := make([]Users, 0)
	_, err := u.DB.Select(&users, fmt.Sprintf("SELECT name, phone, date, city, country, email, coordinates FROM users LIMIT %d", limit))
	if err != nil {
		return nil, err
	}

	return users, nil
}
