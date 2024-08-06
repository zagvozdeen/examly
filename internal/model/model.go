package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type IModel interface {
	GetID() int
}

type Model struct {
	ID        int       `json:"id" db:"id"`
	DeletedAt null.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (m Model) GetID() int {
	return m.ID
}

func (m *Model) FillTime() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}
