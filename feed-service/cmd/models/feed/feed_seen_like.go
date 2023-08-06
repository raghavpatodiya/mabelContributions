package models

type FeedSeenLike struct {
	Username string `json:"username" validate:"required" gorm:"primaryKey"`
	FeedId   string `json:"feed_id" validate:"required" gorm:"primaryKey"`
	Liked    *bool  `json:"liked"`
}
