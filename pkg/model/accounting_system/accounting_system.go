package accounting_system

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	dbconst "github.com/tinkler/erp/pkg/const/db_const"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

type AccountingSystem struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Status      string
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
