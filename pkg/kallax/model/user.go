package model

import "gopkg.in/src-d/go-kallax.v1"

//go:generate kallax gen
type Users struct {
	kallax.Model `table:"users" pk:"id"`
	ID           int64
	Name         string
	Phone        string
	Date         string
	City         string
	Country      string
	Email        string
	Coordinates  string
}
