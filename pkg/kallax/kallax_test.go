package kallax

import (
	"database/sql"
	"strconv"
	"testing"
)

func getUserDB() (*UserDB, error) {
	db, err := sql.Open("postgres", "host=localhost user=docker dbname=test password=dockerpass sslmode=disable")
	if err != nil {
		return nil, err
	}

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
	var i uint64
	for i = 1; i <= 100000; i *= 10 {
		b.Run(strconv.Itoa(int(i)), func(b *testing.B) {
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
