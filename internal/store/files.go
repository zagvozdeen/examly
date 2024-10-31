package store

import (
	"context"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type File struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	Content    string    `json:"content"`
	Size       int       `json:"size"`
	MimeType   string    `json:"mime_type"`
	OriginName string    `json:"origin_name"`
	CreatedBy  int       `json:"created_by"`
	DeletedAt  null.Time `json:"deleted_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type FilesStore interface {
	Create(ctx context.Context, file *File) error
}

type FileStore struct {
	conn *pgxpool.Pool
}

func (s *FileStore) Create(ctx context.Context, file *File) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO files
				(uuid, content, size, mime_type, origin_name, created_by, created_at, updated_at)
			 VALUES
				($1, $2, $3, $4, $5, $6, $7, $8)
			 RETURNING id`,
		file.UUID,
		file.Content,
		file.Size,
		file.MimeType,
		file.OriginName,
		file.CreatedBy,
		file.CreatedAt,
		file.UpdatedAt,
	).Scan(&file.ID)
}
