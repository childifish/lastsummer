package model

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model
	AppKey string `json:"app_key"`
	AppSec string `json:"app_sec"`
}

func (a Auth) Get() {

}

func (a Auth) CheckAuth() {

}
