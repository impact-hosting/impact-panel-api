package models

type APIKey struct {
	APIKey    string `gorm:"primaryKey,index,column:api_key"`
	UserID    uint32 `gorm:"column:user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	Memo      string `gorm:"column:memo,column:memo"`
	CreatedAt int    `gorm:"autoCreatedTime:nano,column:created_at"`
	LastUsed  int    `gorm:"column:last_used"`
}
