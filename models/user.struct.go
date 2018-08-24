package models

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Registered time.Time `json:"registered"`
}

type Users []User