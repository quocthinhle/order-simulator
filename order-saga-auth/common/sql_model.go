package appcommon

import "time"

type SQLModel struct {
	Status    int `gorm:"column:status;default:1;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
