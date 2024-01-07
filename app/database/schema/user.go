package schema

import (
	"github.com/bangadam/go-fiber-starter/utils/helpers"
)

type User struct {
	ID       uint64 `gorm:"primary_key;column:id"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Base
}

// compare password
func (u *User) ComparePassword(password string) bool {
	return helpers.ValidateHash(password, u.Password)
}
