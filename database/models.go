package database

type Task struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Priority int     `json:"priority"`
}
