package HTTPHandlers

import (
	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/8tomat8/yetAnotherCRUD/types"
)

type User struct {
	UserID    int32      `json:"UserID"`
	Username  string     `json:"Username"`
	Password  string     `json:"Password"`
	Firstname string     `json:"Firstname"`
	Lastname  string     `json:"Lastname"`
	Sex       string     `json:"Sex"`
	Birthdate types.Date `json:"Birthdate"`
}

func CreateFromModel(u entity.User) User {
	return User{
		u.UserID,
		u.Username,
		u.Password,
		u.Firstname,
		u.Lastname,
		u.Sex,
		types.Date{Time: u.Birthdate},
	}
}
