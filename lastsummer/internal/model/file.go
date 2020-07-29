package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	FilePath   string `json:"file_path"`   //位置
	Uploader   uint   `json:"uploader"`    //谁上传的
	Now        uint   `json:"now"`         //现在在谁的网盘里
	DangerFile bool   `json:"danger_file"` //没用到
}

func (f *File) WriteFileOrigin(id uint, path string) error {
	f.Uploader = id
	f.FilePath = path
	return DB.Create(&f).Error
}

func (f *File) FindAllFile() (files []File, err error) {
	err = DB.Table("files").Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (f *File) FindMyFile(id uint) (files []File, err error) {
	fmt.Println("id", id)
	err = DB.Table("files").Where("uploader = ?", id).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}
