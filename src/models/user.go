package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uint32    `gorm:"primaryKey,index,column:id"`
	UUID        uuid.UUID `gorm:"column:uuid"`
	Username    string    `gorm:"column:username"`
	Forename    string    `gorm:"column:forename"`
	Surname     string    `gorm:"surname"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:pass"`
	Suspended   bool      `gorm:"column:suspended"`
	Admin       bool      `gorm:"column:admin"`
	MFASecret   string    `gorm:"column:mfa_secret"`
	CreatedAt   int       `gorm:"autoCreatedTime:nano,column:created_at"`
	UpdatedAt   int       `gorm:"autoUpdatedTime:nano,column:updated_at"`
	SuspendedAt int       `gorm:"column:suspended_at"`
}
