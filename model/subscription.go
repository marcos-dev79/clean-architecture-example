package model

type Subscription struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Service string `json:"service"`
	Status  string `json:"status"`
}
