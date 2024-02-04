package models

type Session struct {
	AccessToken  string `gorm:"primaryKey,index,column:access_token"`
	UserID       uint32 `gorm:"column:user_id"`
	User         User   `gorm:"foreignKey:UserID"`
	CreatedAt    int    `gorm:"autoCreateTime:nano,column:created_at"`
	CreatedIP    string `gorm:"column:created_ip"`
	ExpiresAfter int    `gorm:"column:expires_after"`
}
