package schema

import (
	"github.com/bangadam/go-fiber-starter/utils/helpers"
)

type User struct {
	ID        uint64    `gorm:"primary_key;column:id"`
	UserName  *string   `gorm:"column:user_name;null"`
	Password  *string   `gorm:"column:password;null"`
	Base
}

// compare password
func (u *User) ComparePassword(password string) bool {
	return helpers.ValidateHash(password, *u.Password)
}
