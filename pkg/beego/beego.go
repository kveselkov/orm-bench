package beego

import (
	"github.com/astaxie/beego/orm"
	"github.com/kostozyb/orm-bench/internal/config"
)

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}

	// register model
	orm.RegisterModel(new(Users))

	c := config.Config{}

	err = orm.RegisterDataBase("default", c.GetDriver(), c.GetConnectionString(), 30)
	if err != nil {
		panic(err)
	}

	err = orm.RegisterDataBase("test", c.GetDriver(), c.GetConnectionString(), 30)
	if err != nil {
		panic(err)
	}

	err = orm.RunSyncdb("test", false, true)
	if err != nil {
		panic(err)
	}
}

type Users struct {
	Id          int    `orm:"auto"`
	Name        string `orm:"size(255)"`
	Phone       string `orm:"size(100)"`
	Date        string `orm:"size(255)"`
	City        string `orm:"size(255)"`
	Country     string `orm:"size(100)"`
	Email       string `orm:"size(255)"`
	Coordinates string `orm:"size(40)"`
}

type UserDB struct {
	DB orm.Ormer
}

func (u *UserDB) Fetch(limit int) ([]*Users, error) {
	var users []*Users
	qs := u.DB.QueryTable(&Users{})
	_, err := qs.Limit(limit).All(&users)

	return users, err
}
