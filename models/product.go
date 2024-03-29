package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title 		string `json:"title" form:"title" valid:"required~Title is required"`
	Description string `json:"description" form:"description" valid:"required~Description is required"`
	UserId 		uint
	User 		*User
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}
