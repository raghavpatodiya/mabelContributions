package models

type PinBoardShare struct {
	Username string `json:"username" validate:"required" gorm:"primaryKey"`
	BoardId  string `json:"board_id" validate:"required" gorm:"primaryKey"`
}
