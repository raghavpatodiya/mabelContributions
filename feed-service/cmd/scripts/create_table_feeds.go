package scripts

import (
	"super-feed-service/cmd/config"
)

func CreateFeedTable() {
	command := "create table IF NOT EXISTS feeds (" +
		"id text not null primary key," +
		"title TEXT," +
		"asset_code TEXT not null UNIQUE," +
		"link TEXT not null," +
		"thumbnail_link TEXT," +
		"type TEXT not null," +
		"video_length INTEGER," +
		"has_audio BOOLEAN," +
		"asset_quality INTEGER," +
		"asset_ranking INTEGER not null," +
		"product_list jsonb," +
		"brands_list TEXT[]," +
		"category_list TEXT[]," +
		"usecase_list TEXT[]," +
		"tags TEXT[]," +
		"product_type_list TEXT[]," +
		"live Boolean NOT NULL default false," +
		"created_at TIMESTAMP NOT NULL default now()," +
		"updated_at TIMESTAMP," +
		"is_deleted boolean default false," +
		"is_updated boolean default false," +
		"likes_count Numeric NOT NULL default 0," +
		"shares_count Numeric NOT NULL default 0," +
		"dimension TEXT NOT NULL" +
		");"
	err := config.DB.Exec(command).Error
	if err != nil {
		panic(err)
	}
}
