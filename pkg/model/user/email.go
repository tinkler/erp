package user

import (
	dbconst "github.com/tinkler/erp/pkg/const/db_const"
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	UserID  string `gorm:"type:uuid"`
	Address string `gorm:"type:varchar(255);not null"`
}

func (Email) TableName() string { return dbconst.Schema + ".emails" }
