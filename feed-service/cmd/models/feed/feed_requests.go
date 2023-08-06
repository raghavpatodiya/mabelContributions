package models

type RequestBody struct {
	Title             string   `json:"title"`
	Type              string   `json:"type"`
	Link              string   `json:"url"`
	AssetCode         string   `json:"asset_code"`
	Skuslist          []string `json:"product_list"`
	VideoLength       int64    `json:"video_length"`
	HasAudio          bool     `json:"has_audio"`
	AssetQuality      int      `json:"asset_quality"`
	AssetRanking      int      `json:"asset_ranking"`
	Brands_List       []string `json:"brands_list"`
	Category_List     []string `json:"category_list"`
	Usecase_List      []string `json:"usecase_list"`
	Product_Type_List []string `json:"product_type_list"`
	Live              bool     `json:"live"`
	Tags              []string `json:"tags"`
}
