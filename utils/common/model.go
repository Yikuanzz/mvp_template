package common

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`              // 创建时间，自动设置为当前时间
	UpdatedAt time.Time      `json:"updated_at"`              // 更新时间，自动更新为当前时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // 删除时间，索引字段
}
