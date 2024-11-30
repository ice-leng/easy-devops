package model

type RoleMenu struct {
	Model `gorm:"embedded"` // embed id and time

	RoleID uint64 `gorm:"column:role_id;type:int(11);NOT NULL" json:"roleID"`
	MenuID uint64 `gorm:"column:menu_id;type:int(11);NOT NULL" json:"menuID"`
}

// TableName table name
func (m *RoleMenu) TableName() string {
	return "t_role_menu"
}
