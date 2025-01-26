package model

// Activity struct
type Activity struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

type Repo struct {
	Name string `json:"name"`
}
