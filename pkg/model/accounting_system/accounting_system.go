package accounting_system

import (
	"context"
	"net/http"
	"os/user"

	"github.com/google/uuid"
	dbconst "github.com/tinkler/erp/pkg/const/db_const"
	"github.com/tinkler/erp/pkg/model/company"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

type AccountingSystem struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Status      string
	Users       []*user.User       `gorm:"many2many:user_accounting_systems"`
	Company     []*company.Company `gorm:"many2many:company_accounting_systems"`
}

func (AccountingSystem) TableName() string {
	return dbconst.Schema + ".accounting_systems"
}

func (m AccountingSystem) ValidateID() error {
	if _, err := uuid.Parse(m.ID); err != nil {
		return status.NewCn(http.StatusBadRequest, "Invalid ID", "ID非法")
	}
	return nil
}

func (m *AccountingSystem) Create(ctx context.Context) error {
	return nil
}
