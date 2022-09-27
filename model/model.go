package model

type Todo struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Status int    `json:"status"`
}
