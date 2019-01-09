package kallax

import (
	"database/sql"

	"github.com/kostozyb/orm-bench/pkg/kallax/model"
)

type UserDB struct {
	DB *sql.DB
}

func (u *UserDB) Fetch(limit uint64) ([]model.Users, error) {
	store := model.NewUsersStore(u.DB)
	query := model.NewUsersQuery().Limit(limit)
	us, err := store.Find(query)
	if err != nil {
		return nil, err
	}

	var users []model.Users
	for us.Next() {
		u, err := us.Get()
		if err != nil {
			return nil, err
		}

		users = append(users, *u)
	}

	return users, nil
}
