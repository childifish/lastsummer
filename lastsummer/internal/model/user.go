package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Level    int    `json:"level"`
}

func (u *User) GetUid() uint {
	DB.Where("username = ?", u.Username).First(&u)
	return u.ID
}

func (u *User) Write() error {
	return DB.Create(&u).Error
}

func (u *User) CheckRegister(username string) bool {
	DB.Where("username = ?", username).First(&u)
	if u.ID != 0 {
		return false
	}
	return true
}

func (u *User) CheckLogin() bool {
	DB.Where(User{
		Username: u.Username,
		Password: u.Password,
	}).First(&u)
	if u.ID != 0 {
		return true
	}
	return false
}
