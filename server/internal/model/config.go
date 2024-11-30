package model

// Config 系统配置
type Config struct {
	Model `gorm:"embedded"` // embed id and time

	Name        string `gorm:"column:name;type:text;NOT NULL" json:"name"`               // 配置名称
	Description string `gorm:"column:description;type:text;NOT NULL" json:"description"` // 描述
	Key         string `gorm:"column:key;type:text;NOT NULL" json:"key"`                 // 配置键
	Value       string `gorm:"column:value;type:text;NOT NULL" json:"value"`             // 配置值
}

// TableName table name
func (m *Config) TableName() string {
	return "t_config"
}
