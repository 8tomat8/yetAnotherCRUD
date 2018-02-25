package entity

import (
	"crypto/md5"

	"encoding/hex"

	"time"

	"github.com/pkg/errors"
)

const passwordHashSalt = "KAuyg3ofbacvn8euvbtcvkwtyvqwue368q&^#Q"

// User represent user struct
type User struct {
	UserID    int32
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Sex       string
	Birthdate time.Time
}

func (u User) IsValid() (valid bool) {
	switch {
	case u.Birthdate == time.Time{}:
		return
	case u.Username == "":
		return
	case u.Password == "":
		return
	case u.Firstname == "":
		return
	case u.Lastname == "":
		return
	case u.Sex == "":
		return
	}

	return true
}

// SetPassword calculates hash for new password and saves it into Password
func (u *User) SetPassword(plainPassword string) error {
	hashedPass, err := calculateUserHash(plainPassword)
	if err != nil {
		return errors.Wrap(err, "cannot calculate hash from password")
	}
	u.Password = hashedPass
	return nil
}

func calculateUserHash(plainPassword string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(passwordHashSalt + plainPassword + passwordHashSalt))
	if err != nil {
		return "", errors.Wrap(err, "cannot write to hash object")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
