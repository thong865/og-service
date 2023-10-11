package models

import (
	hash "MS/config/hash"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255" json:"username"`
	MobileNo string `gorm:"size:255" json:"mobile_no"`
	Password string `gorm:"size:255" json:"password"`
}

func (u *User) TableName() string {
	return "sstm_users"
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	p, err := hash.Encrypt(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(p)

	return nil

}
