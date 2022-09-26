package model

type Todo struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Status string `json:"status"`
}
