package models

type PinFeedContent struct {
	FeedID  string `json:"feed_id" validate:"required"`
	BoardID string `json:"board_id" validate:"required"`
}
