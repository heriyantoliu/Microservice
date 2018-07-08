package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid2 "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, _ := uuid2.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
