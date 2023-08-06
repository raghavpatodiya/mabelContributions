package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

// type Feed struct {
// 	Id                string         `json:"id"`
// 	Title             string         `json:"title" validate:"required"`
// 	AssetCode         string         `json:"asset_code"`
// 	Link              string         `json:"link" validate:"required"`
// 	Thumbnail_Link    string         `json:"thumbnail_link" validate:"required"`
// 	Type              string         `json:"type"`
// 	Product_List      pq.StringArray `gorm:"type:text[]"`
// 	VideoLength       int64          `json:"video_length"`
// 	HasAudio          bool           `json:"has_audio"`
// 	AssetQuality      int            `json:"asset_quality"`
// 	AssetRanking      int            `json:"asset_ranking"`
// 	Brands_List       pq.StringArray `gorm:"type:text[]"`
// 	Category_List     pq.StringArray `gorm:"type:text[]"`
// 	Usecase_List      pq.StringArray `gorm:"type:text[]"`
// 	Product_Type_List pq.StringArray `gorm:"type:text[]"`
// 	Live              bool           `json:"live"`
// 	Tags              pq.StringArray `gorm:"type:text[]"`
// 	CreatedAt         time.Time      `gorm:"created_at"`
// 	UpdatedAt         time.Time      `gorm:"updated_at"`
// 	Is_Deleted        bool           `json:"is_deleted"`
// 	LikesCount        int64          `json:"likes_count"`
// 	SharesCount       int64          `json:"share_count"`
// }

// type Feed_Cluster_Map struct {
// 	Id         string `json:"id"`
// 	Feed_Id    string `json:"feed_id"`
// 	Cluster_Id string `json:"cluster_id"`
// }
type Feed struct {
	Id                string          `json:"id"`
	Title             string          `json:"title" validate:"required"`
	AssetCode         string          `json:"asset_code"`
	Link              string          `json:"link" validate:"required"`
	Thumbnail_Link    string          `json:"thumbnail_link" validate:"required"`
	Type              string          `json:"type"`
	Product_List      json.RawMessage `json:"product_list"`
	VideoLength       int64           `json:"video_length"`
	HasAudio          bool            `json:"has_audio"`
	AssetQuality      int             `json:"asset_quality"`
	AssetRanking      int             `json:"asset_ranking"`
	Brands_List       pq.StringArray  `gorm:"type:text[]"`
	Category_List     pq.StringArray  `gorm:"type:text[]"`
	Usecase_List      pq.StringArray  `gorm:"type:text[]"`
	Product_Type_List pq.StringArray  `gorm:"type:text[]"`
	Live              bool            `json:"live"`
	Tags              pq.StringArray  `gorm:"type:text[]"`
	CreatedAt         time.Time       `gorm:"created_at"`
	UpdatedAt         time.Time       `gorm:"updated_at"`
	Is_Deleted        bool            `json:"is_deleted"`
	Is_Updated        bool            `json:"is_updated"`
	LikesCount        int64           `json:"likes_count"`
	SharesCount       int64           `json:"share_count"`
	Dimension         string          `json:"dimension"`
}

// type FeedStructure interface {
// 	Save()
// 	Update()
// 	Delete()
// 	FindAll()
// }
// type database struct {
// 	connection *gorm.DB
// }

// func NewFeedStructure() FeedStructure {
// 	db, err := gorm.Open("postgres", "mabel_feed.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.AutoMigrate()
// }

type Feed_Cluster_Map struct {
	Id         string `gorm:"primary_key;auto_increment" json:"id"`
	Feed_Id    string `gorm:"type:varchar()" json:"feed_id"`
	Cluster_Id string `gorm:"type:varchar()" json:"cluster_id"`
}
