package user

import (
	dbconst "github.com/tinkler/erp/pkg/const/db_const"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	UserID   string `gorm:"type:uuid"`
	Password string `gorm:"type:varchar(64);not null"`
}

func (Auth) TableName() string {
	return dbconst.Schema + ".auths"
}
