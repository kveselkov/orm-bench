package sqlboiler

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/kostozyb/orm-bench/pkg/sqlboiler/models"
)

//go:generate sqlboiler --wipe psql
type Users struct {
	ID          int32  `boil:"id"`
	Name        string `boil:"name"`
	Phone       string `boil:"phone"`
	Date        string `boil:"date"`
	City        string `boil:"city"`
	Country     string `boil:"country"`
	Email       string `boil:"email"`
	Coordinates string `boil:"coordinates"`
}

type UserDB struct {
	DB *sql.DB
}

func (u *UserDB) Fetch(limit int) (models.UserSlice, error) {
	return models.Users(
		qm.Limit(limit),
	).All(context.Background(), u.DB)
}
