package model

type Role struct {
	Model `gorm:"embedded"` // embed id and time

	Name   string `gorm:"column:name;type:text;NOT NULL" json:"name"`
	Code   string `gorm:"column:code;type:text;NOT NULL" json:"code"`
	Sort   int    `gorm:"column:sort;type:int(11);NOT NULL" json:"sort"`
	Status int    `gorm:"column:status;type:int(11);NOT NULL" json:"status"`
}

// TableName table name
func (m *Role) TableName() string {
	return "t_role"
}
