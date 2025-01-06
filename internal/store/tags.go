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

//func (s *TagStore) UpdateQuestionTags(ctx context.Context, questionId int, tagsIds []int) (err error) {
//	args := make([]any, len(tagsIds)+1)
//	args[0] = questionId
//	placeholders := make([]string, len(tagsIds))
//	for i, id := range tagsIds {
//		_, err = s.conn.Exec(
//			ctx,
//			"INSERT INTO question_tag (question_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
//			questionId,
//			id,
//		)
//		if err != nil {
//			return err
//		}
//		placeholders[i] = fmt.Sprintf("$%d", i+2)
//		args[i+1] = id
//	}
//	placeholder := strings.Join(placeholders, ",")
//	_, err = s.conn.Exec(
//		ctx,
//		fmt.Sprintf("DELETE FROM question_tag WHERE question_id = $1 AND tag_id NOT IN (%s)", placeholder),
//		args...,
//	)
//	return err
//}

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
