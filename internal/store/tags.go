package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagsStore interface {
	Get(ctx context.Context) ([]Tag, error)
	CreateQuestionTags(ctx context.Context, questionId int, ids []int) (err error)
	GetQuestionTags(ctx context.Context, id int) ([]int, error)
	GetTagsQuestions(ctx context.Context, tagsIds []int) ([]int, error)
}

type TagStore struct {
	conn *pgxpool.Pool
}

func (s *TagStore) Get(ctx context.Context) (tags []Tag, err error) {
	rows, err := s.conn.Query(ctx, "SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		err = rows.Scan(
			&tag.ID,
			&tag.Name,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *TagStore) CreateQuestionTags(ctx context.Context, questionId int, ids []int) (err error) {
	args := make([]any, len(ids)+1)
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("($1, $%d)", i+2)
		args[i+1] = id
	}
	placeholder := strings.Join(placeholders, ",")
	args[0] = questionId
	_, err = s.conn.Exec(
		ctx,
		fmt.Sprintf("INSERT INTO question_tag (question_id, tag_id) VALUES %s", placeholder),
		args...,
	)
	return err
}

func (s *TagStore) GetQuestionTags(ctx context.Context, id int) (ids []int, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT tag_id FROM question_tag WHERE question_id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (s *TagStore) GetTagsQuestions(ctx context.Context, tagsIds []int) (ids []int, err error) {
	placeholders := make([]string, len(tagsIds))
	args := make([]any, len(tagsIds))
	for i, id := range tagsIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}
	sql := fmt.Sprintf("SELECT DISTINCT question_id FROM question_tag WHERE tag_id IN (%s)", strings.Join(placeholders, ","))
	rows, err := s.conn.Query(ctx, sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids, id)
	}
	return
}
