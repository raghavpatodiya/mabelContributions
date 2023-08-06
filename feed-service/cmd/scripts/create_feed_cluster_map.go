package scripts

import (
	"super-feed-service/cmd/config"
)

func CreateFeedClusterMapTable(){
	command:= "CREATE TABLE feed_cluster_map ("+
    "id text,"+
    "cluster_id text,"+
    "feed_id text,"+
    "PRIMARY KEY (id)"+
	");";
	err:= config.DB.Exec(command).Error
	if err != nil {
		panic(err)
	}
}