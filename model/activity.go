package model

// Activity 活动结构体
type Activity struct {
	ID    string `json:"id"` // 修改类型为 string
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
	Repo  Repo   `json:"repo"`
}

// Actor 执行者结构体
type Actor struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

// Repo 仓库结构体
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
