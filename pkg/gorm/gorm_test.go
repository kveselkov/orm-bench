package gorm

import (
	"github.com/kostozyb/orm-bench/internal/config"
	"strconv"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getUserDB(driver, cs string) (*UserDB, error) {
	db, err := gorm.Open(driver, cs)
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
			c := config.Config{}

			u, err := getUserDB(c.GetDriver(), c.GetConnectionString())
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
