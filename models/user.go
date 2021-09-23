package models

import (
	"github.com/asaskevich/govalidator"
	"go-gin-jwt/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName	string `gorm:"not null" valid:"required~Full Name is required" json:"full_name" form:"full_name"`
	Email		string `gorm:"not null;uniqueIndex" valid:"required~Email is required,email~Invalid email format" json:"email" form:"email"`
	Password	string `gorm:"not null" valid:"required~Password is required,minstringlength(5)~Password must have 5 character or more" json:"password" form:"password"`
	Products	[]Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.HashPassword(u.Password)
	return nil
}