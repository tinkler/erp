package commodity

import (
	"context"
	"strings"

	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
)

type CommodityCollection struct {
	List []*Commodity
}

func (c *CommodityCollection) SearchByName(ctx context.Context, namePattern string) error {
	namePattern = strings.TrimSpace(namePattern)
	if len(namePattern) == 0 {
		return status.StatusBadRequest("empty search pattern")
	}
	if err := db.GetDB(ctx).Model(&Commodity{}).Where("name LIKE ?", namePattern).Find(&c.List).Error; err != nil {
		return status.StatusInternalServer(err)
	}
	return nil
}
func (c *CommodityCollection) SearchByCode(ctx context.Context, codePattern string) error {
	codePattern = strings.TrimSpace(codePattern)
	if len(codePattern) == 0 {
		return status.StatusBadRequest("empty search pattern")
	}
	if err := db.GetDB(ctx).Model(&Commodity{}).Where("barcode LIKE ?", codePattern).Find(&c.List).Error; err != nil {
		return status.StatusInternalServer(err)
	}
	return nil
}
