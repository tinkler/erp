package commodity

import (
	"context"
	"net/http"
	"strings"

	"github.com/tinkler/erp/mist/pssm"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

type Commodity struct {
	gorm.Model
	ID       string `gorm:"primaryKey;autoIncrement;"`
	Name     string `gorm:"unique"`
	Price    float64
	Stock    int
	Category string
	ImageUrl string
	Unit     string
	Barcode  string
}

func (m *Commodity) Create(ctx context.Context) error {
	if m.ID != "" {
		return status.NewCn(http.StatusBadRequest, "ID is not empty", "新增商品ID不为空")
	}
	m.Name = strings.TrimSpace(m.Name)
	if len(m.Name) == 0 {
		return status.NewCn(http.StatusBadRequest, "Name is empty", "新增商品名称为空")
	}
	m.Barcode = strings.TrimSpace(m.Barcode)
	if len(m.Barcode) == 0 {
		return status.NewCn(http.StatusBadRequest, "Barcode is empty", "新增商品二维码为空")
	}
	m.Unit = strings.TrimSpace(m.Unit)
	if len(m.Unit) == 0 {
		return status.NewCn(http.StatusBadRequest, "Unit is empty", "新增商品单位为空")
	}
	m.Category = strings.TrimSpace(m.Category)
	if len(m.Category) == 0 {
		return status.NewCn(http.StatusBadRequest, "Category is empty", "新增商品类别为空")
	}
	if err := pssm.DB(ctx).Create(m).Error; err != nil {
		return status.StatusInternalServer(err)
	}

	return nil
}
