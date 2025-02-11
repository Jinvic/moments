package db

import (
	"time"
)

type Notice struct {
	Id        int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	UserId    int32      `gorm:"column:userId;NOT NULL" json:"userId,omitempty"`
	MemoId    int32      `gorm:"column:memoId;NOT NULL" json:"memoId,omitempty"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	Type      int32      `gorm:"column:type;NOT NULL" json:"type,omitempty"`
	IsRead    bool       `gorm:"column:isRead;NOT NULL" json:"isRead,omitempty"`
	AvatarUrl string     `gorm:"column:avatarUrl" json:"avatarUrl,omitempty"`
	From      string     `gorm:"column:from" json:"from,omitempty"`
	Content   string     `gorm:"column:content" json:"content,omitempty"`
}

func (n *Notice) TableName() string {
	return "Notice"
}
