package models

import (
	"time"
)

type PinBoard struct {
	BoardID     string     `json:"board_id" validate:"required" gorm:"primaryKey"`
	Username    string     `json:"username"`
	BoardName   string     `json:"board_name" validate:"required"`
	IsPublic    bool       `json:"is_public"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	IsUpdated   bool       `json:"is_updated"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	LikesCount  int64      `json:"likes_count"`
	SharesCount int64      `json:"share_count"`
}
