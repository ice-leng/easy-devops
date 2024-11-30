package model

import (
	"admin/internal/types"
)

// Model embedded structs, add `gorm: "embedded"` when defining table structs
type Model struct {
	ID        uint64              `gorm:"column:id" json:"id"`
	CreatedAt types.LocalDateTime `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt types.LocalDateTime `gorm:"column:updated_at" json:"updatedAt"`
}
