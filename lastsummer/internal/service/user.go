package service

import (
	"lastsummer/internal/model"
)

func Register(username string, password string) (bool, string) {
	var u model.User
	if !u.CheckRegister(username) {
		return false, "重复用户名"
	}
	u.Username = username
	u.Password = password
	if err := u.Write(); err != nil {
		//log.Print(err)
		return false, "写入错误"
	}
	return true, ""
}

func Login(username string, password string) bool {
	var u model.User
	u.Username = username
	u.Password = password
	if ok := u.CheckLogin(); ok {
		return true
	}
	return false
}

func Vip() {

}
