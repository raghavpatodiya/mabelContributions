package models

type User struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ClusterId string `json:"cluster_id"`
	CreatedAt string `json:"create_at"`
}
