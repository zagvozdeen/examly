package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) CreateFile(file *model.File) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO files (uuid, content, size, mime_type, origin_name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		file.UUID,
		file.Content,
		file.Size,
		file.MimeType,
		file.OriginName,
		file.UserID,
		file.CreatedAt,
		file.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *FileRepository) GetFileByID(id int) (file model.File, err error) {
	err = r.db.Get(
		&file,
		"SELECT * FROM files WHERE id = $1 AND deleted_at IS NULL",
		id,
	)

	return file, err
}
