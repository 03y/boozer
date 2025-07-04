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

type LeaderboardItem struct {
	Consumed int `json:"consumed"`
	Item
}
