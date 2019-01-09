package sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	DB *sqlx.DB
}

func (u *UserDB) Fetch(limit int) ([]Users, error) {
	users := []Users{}
	err := u.DB.Select(&users, fmt.Sprintf("SELECT * FROM users LIMIT %d", limit))
	if err != nil {
		return nil, err
	}

	return users, nil
}
