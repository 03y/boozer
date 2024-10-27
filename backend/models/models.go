package models

type User struct {
	User_id  int
	Username string
	Password string // hashed, never in clear
	Created  int    // unix timestamp
}

type Item struct {
	Item_id int
	Name    string
	Units   float32
	Added   int // unix timestamp
}

type Consumption struct {
	Consumption_id int
	User_id        int
	Item_id        int
	Time           int // unix timestamp
}

type LeaderboardUser struct {
	Username string
	Consumed int
}

type LeaderboardItem struct {
	Consumed int
	Item
}
