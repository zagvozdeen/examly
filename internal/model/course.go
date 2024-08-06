package model

type Course struct {
	Model
	UUID        string `json:"uuid" db:"uuid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Color       string `json:"color" db:"color"`
	Icon        string `json:"icon" db:"icon"`
	Status      string `json:"status" db:"status"`
	UserID      int    `json:"user_id" db:"user_id"`
}
