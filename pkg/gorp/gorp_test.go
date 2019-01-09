package gorp

import (
	"database/sql"
	"strconv"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gorp.v1"
)

func getUserDB() (*UserDB, error) {
	db, err := sql.Open("postgres", "host=localhost user=docker dbname=test password=dockerpass sslmode=disable")
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	u := &UserDB{
		DB: dbMap,
	}

	return u, nil
}

func BenchmarkUserDB_Fetch(b *testing.B) {
	for i := 1; i <= 100000; i *= 10 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			u, err := getUserDB()
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				_, err = u.Fetch(i)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
