package model

import "time"

type BaseColumn struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:创建者" json:"updatedAt"`
	CreatedBy string    `gorm:"column:created_by;comment:创建者" json:"createBy"`
	UpdatedBy string    `gorm:"column:updated_by;comment:创建者" json:"updatedBy"`
}

type BaseColumnWithStatus struct {
	BaseColumn
	Status uint `gorm:"type:smallint;default:1;comment:'1正常, 2禁用'" json:"status"`
}

// type BaseColumnWithPermissions struct {
//   BaseColumn
// }
