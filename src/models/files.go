package models

import "gorm.io/gorm"

type FileBase struct {
	Name string `json:"name"`
	Size uint   `json:"size"`
	Ext  string `json:"ext"`
}

type File struct {
	FileBase
	gorm.Model
}
