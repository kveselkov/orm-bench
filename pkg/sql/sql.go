package sql

import (
	"database/sql"
	"fmt"
	"log"
)

type Users struct {
	ID          int32
	Name        string
	Phone       string
	Date        string
	City        string
	Country     string
	Email       string
	Coordinates string
}

type UserDB struct {
	DB *sql.DB
}

func (u *UserDB) Fetch(limit int) ([]*Users, error) {
	s := fmt.Sprintf("SELECT id, name, phone, date, city, country, email, coordinates FROM users LIMIT %d", limit)
	rows, err := u.DB.Query(s)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Printf("error close rows: %s\n", err)
		}
	}()

	users := make([]*Users, 0)
	for rows.Next() {
		u := &Users{}
		err := rows.Scan(&u.ID, &u.Name, &u.Phone, &u.Date, &u.City, &u.Country, &u.Email, &u.Coordinates)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
