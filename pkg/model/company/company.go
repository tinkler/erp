package company

import (
	"github.com/tinkler/erp/pkg/model/user"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID    string `gorm:"primaryKey;autoIncrement;"`
	Name  string
	Users []*user.User `gorm:"many2many:company_users;"`
}
