package storage

import "time"

type userModel struct {
	UserID    int32     `json:"UserID"`
	Username  string    `json:"Username"`
	Password  string    `json:"Password"`
	Firstname string    `json:"Firstname"`
	Lastname  string    `json:"Lastname"`
	Sex       string    `json:"Sex"`
	Birthdate time.Time `json:"Birthdate"`
}
