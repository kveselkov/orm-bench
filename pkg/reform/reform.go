package reform

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
)

//go:generate reform

// Users represents a row in users table.
//reform:users
type Users struct {
	ID          int32  `reform:"id,pk"`
	Name        string `reform:"name"`
	Phone       string `reform:"phone"`
	Date        string `reform:"date"`
	City        string `reform:"city"`
	Country     string `reform:"country"`
	Email       string `reform:"email"`
	Coordinates string `reform:"coordinates"`
}

type UserDB struct {
	DB *reform.DB
}

func (u *UserDB) Fetch(limit int) ([]Users, error) {
	tail := fmt.Sprintf("LIMIT %d", limit)
	rows, err := u.DB.SelectRows(UsersTable, tail)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Printf("Error close rows: %s", err)
		}
	}()

	users := make([]Users, 0)
	for {
		var user Users
		if err = u.DB.NextRow(&user, rows); err != nil {
			break
		}
		users = append(users, user)
	}

	if err != reform.ErrNoRows {
		return nil, err
	}

	return users, nil
}
