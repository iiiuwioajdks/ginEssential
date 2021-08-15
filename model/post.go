package model

import uuid "github.com/satori/go.uuid"
import "gorm.io/gorm"

/**
文章
*/

type Post struct {
	ID     uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId uint      `json:"user_id" gorm:"not null"`
	// 这个 CategoryId 充当 Category 的外键
	CategoryId uint `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}

// 在文章创建前赋值 uuid

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.NewV4()
	return nil
}
