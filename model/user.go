package model

type User struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"primaryKey"`
	Body   string
	Number string
}
