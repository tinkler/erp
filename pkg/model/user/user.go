package user

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	dbconst "github.com/tinkler/erp/pkg/const/db_const"
	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string   `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Username    string   `gorm:"type:varchar(100);"`
	PhoneNumber string   `gorm:"type:varchar(20);"`
	Emails      []*Email `gorm:"foreignKey:UserID"`
	Auth        *Auth    `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return dbconst.Schema + ".users"
}

func (m User) ValidateID() error {
	if _, err := uuid.Parse(m.ID); err != nil {
		return status.NewCn(http.StatusBadRequest, "ID invalid", "ID非法")
	}
	return nil
}

func (m *User) Filter() error {
	m.Username = strings.TrimSpace(m.Username)
	return nil
}

func (m *User) Get(ctx context.Context) error {
	if err := m.Filter(); err != nil {
		return err
	}
	idErr := m.ValidateID()
	if idErr != nil {
		if len(m.Username) == 0 {
			return idErr
		}
	}
	return db.DB().WithContext(ctx).First(m).Error
}

func (m *User) GetDetail(ctx context.Context) error {
	if err := m.ValidateID(); err != nil {
		return err
	}
	se := db.DB().WithContext(ctx)
	if err := se.Model(&User{ID: m.ID}).Association("Emails").Find(&m.Emails); err != nil {
		return status.StatusInternalServer(err)
	}
	return nil
}
