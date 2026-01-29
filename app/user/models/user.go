package models

import (
	"go-admin/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"id"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	models.ControlBy
	models.ModelTime
}

func (*User) TableName() string {
	return "user"
}

func (e *User) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *User) GetId() interface{} {
	return e.Id
}

// Encrypt 加密
func (e *User) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *User) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *User) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}
