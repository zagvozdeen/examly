package model

import "time"

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

type FullCourseStats struct {
	ID            int       `json:"id" db:"id"`
	UUID          string    `json:"uuid" db:"uuid"`
	Type          string    `json:"type" db:"type"`
	CreatedAt     time.Time `json:"-" db:"created_at"`
	CreatedAtUnix int64     `json:"created_at" db:"-"`
	Correct       int       `json:"correct" db:"correct"`
	Incorrect     int       `json:"incorrect" db:"incorrect"`
	Total         int       `json:"total" db:"total"`
}

type CourseStatsParams struct {
	CourseID int
	UserID   int
}
