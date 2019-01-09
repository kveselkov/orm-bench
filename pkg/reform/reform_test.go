package reform

import (
	"database/sql"
	"strconv"
	"testing"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type EmptyLogger struct {
}

func (EmptyLogger) Printf(format string, v ...interface{}) {

}

func getUserDB() (*UserDB, error) {
	conn, err := sql.Open("postgres", "host=localhost user=docker dbname=test password=dockerpass sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	logger := &EmptyLogger{}

	db := reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

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
