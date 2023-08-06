package models

type FeedShare struct {
	Username string `json:"username" validate:"required" gorm:"primaryKey"`
	FeedId   string `json:"feed_id" validate:"required" gorm:"primaryKey"`
}
