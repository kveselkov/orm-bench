package xorm

import (
	"strconv"
	"testing"

	"github.com/go-xorm/core"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getUserDB() (*UserDB, error) {
	db, err := xorm.NewEngine("postgres", "host=localhost user=docker dbname=test password=dockerpass sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.Logger().SetLevel(core.LOG_OFF)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	u := &UserDB{
		DB: db,
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
