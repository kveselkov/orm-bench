package beego

import (
	"strconv"
	"testing"

	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func getUserDB() *UserDB {
	o := orm.NewOrm()

	return &UserDB{
		DB: o,
	}
}

func BenchmarkUserDB_Fetch(b *testing.B) {
	for i := 1; i <= 100000; i *= 10 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			u := getUserDB()

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				_, err := u.Fetch(i)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
