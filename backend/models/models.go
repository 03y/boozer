package models

type User struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"` // hashed, never in clear
	Created  int    `json:"created"`  // unix timestamp
}

type UserNoPw struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	Created  int    `json:"created"` // unix timestamp
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type Item struct {
	Item_id int     `json:"item_id"`
	Name    string  `json:"name"`
	Units   float32 `json:"units"`
	Added   int     `json:"added"` // unix timestamp
}

type Consumption struct {
	Consumption_id int `json:"consumption_id"`
	User_id        int `json:"user_id"`
	Item_id        int `json:"item_id"`
	Time           int `json:"time"` // unix timestamp
}

type LeaderboardUser struct {
	Consumed int    `json:"consumed"`
	Username string `json:"username"`
}

type LeaderboardUserUnits struct {
	Units    float32 `json:"units"`
	Username string  `json:"username"`
}

type LeaderboardItem struct {
	Consumed int `json:"consumed"`
	Item
}

type FeedConsumption struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Time     int    `json:"time"` // unix timestamp
}
