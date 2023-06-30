package inventory

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/tinkler/erp"
	"github.com/tinkler/erp/mist/pssm/pkg/model/commodity"
	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

type StoreIn struct {
	gorm.Model
	ID           string         `gorm:"primaryKey;autoIncrement;"`
	UserID       string         `gorm:"type:uuid"`
	StoreInItems []*StoreInItem `gorm:"foreignKey:StoreInID"`
}

type StoreInItem struct {
	gorm.Model
	ID          uint `gorm:"primarykey;autoIncrement"`
	Index       int
	StoreInID   string `gorm:"type:uuid"`
	CommodityID string `gorm:"type:uuid"`
	Price       int
	Quantity    float64
}

func (m *StoreInItem) Validate() error {
	if m.Quantity == 0 {
		return status.StatusBadRequest("Quantity must be greater than zero")
	}
	if _, err := uuid.Parse(m.CommodityID); err != nil {
		return status.StatusBadRequest("CommodityID is invalid")
	}
	return nil
}

func (m *StoreIn) Create(ctx context.Context) error {
	m.ID = strings.TrimSpace(m.ID)
	if m.ID != "" {
		return status.StatusBadRequest("ID must be empty")
	}
	for i := range m.StoreInItems {
		err := m.StoreInItems[i].Validate()
		if err != nil {
			return err
		}
	}
	m.sortItems()

	m.UserID = erp.GetLoginUser(ctx).ID()
	tx := db.GetDB(ctx).Begin()
	if tx.Error != nil {
		return status.StatusInternalServer(tx.Error)
	}

	for i := range m.StoreInItems {
		if err := tx.Model(&commodity.Commodity{}).Update("stock", gorm.Expr("stock + ?", m.StoreInItems[i].Quantity)).Error; err != nil {
			_ = tx.Rollback()
			return status.StatusInternalServer(err)
		}
	}

	if err := tx.Create(m).Error; err != nil {
		_ = tx.Rollback()
		return status.StatusInternalServer(err)
	}

	for i := range m.StoreInItems {
		m.StoreInItems[i].StoreInID = m.ID
	}
	if err := tx.CreateInBatches(m.StoreInItems, len(m.StoreInItems)).Error; err != nil {
		_ = tx.Rollback()
		return status.StatusInternalServer(err)
	}
	if err := tx.Commit().Error; err != nil {
		return status.StatusInternalServer(err)
	}

	return nil
}
