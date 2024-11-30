package model

import (
	"admin/internal/types"
)

type Platform struct {
	Model `gorm:"embedded"` // embed id and time

	Username string               `gorm:"column:username;type:text;NOT NULL" json:"username"`
	Password string               `gorm:"column:password;type:text;NOT NULL" json:"password"`
	Avatar   string               `gorm:"column:avatar;type:text;NOT NULL" json:"avatar"`
	RoleID   types.LocalIntArray  `gorm:"column:role_id;type:text;NOT NULL" json:"roleID"`
	Status   *int                 `gorm:"column:status;type:int(11);NOT NULL" json:"status"`
	LastTime *types.LocalDateTime `gorm:"column:last_time;type:text" json:"lastTime"`
}

// TableName table name
func (m *Platform) TableName() string {
	return "t_platform"
}
