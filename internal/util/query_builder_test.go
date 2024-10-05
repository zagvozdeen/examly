package util

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	t.Run("should build a simple query", func(t *testing.T) {
		b := NewQueryBuilder("SELECT * FROM users").
			WhereNull("deleted_at").
			WhereFunc(func(b *QueryBuilder) {
				b.Where("user_id", "=", 2).
					OrWhere("status", "=", "active")
			})

		assert.Equal(
			t,
			"SELECT * FROM users WHERE deleted_at IS NULL AND (user_id = $1 OR status = $2)",
			b.String(),
		)
		assert.Equal(t, []any{2, "active"}, b.params)
	})

	t.Run("should build a very simple query", func(t *testing.T) {
		b := NewQueryBuilder("SELECT * FROM users").WhereNull("deleted_at")

		assert.Equal(
			t,
			"SELECT * FROM users WHERE deleted_at IS NULL",
			b.String(),
		)
		assert.Equal(t, nil, b.params)
	})

	t.Run("should build simple or query", func(t *testing.T) {
		b := NewQueryBuilder("SELECT * FROM users").
			WhereFunc(func(b *QueryBuilder) {
				b.Where("user_id", "=", 2).
					OrWhere("status", "=", "active")
			})

		assert.Equal(
			t,
			"SELECT * FROM users WHERE (user_id = $1 OR status = $2)",
			b.String(),
		)
		assert.Equal(t, []any{2, "active"}, b.params)
	})

	t.Run("should build medium query", func(t *testing.T) {
		b := NewQueryBuilder("SELECT * FROM users").
			WhereFunc(func(b *QueryBuilder) {
				b.WhereFunc(func(b *QueryBuilder) {
					b.Where("user_id", "=", 4).Where("test", "=", 1)
				}).OrWhereFunc(func(b *QueryBuilder) {
					b.Where("user_id", "=", 5).Where("test", "!=", 5)
				})
			})

		assert.Equal(
			t,
			"SELECT * FROM users WHERE ((user_id = $1 AND test = $2) OR (user_id = $3 AND test != $4))",
			b.String(),
		)
		assert.Equal(t, []any{4, 1, 5, 5}, b.params)
	})

	t.Run("should build where in query", func(t *testing.T) {
		b := NewQueryBuilder("SELECT * FROM users").
			Where("status", "=", "active").
			WhereIn("user_id", []any{1, 2, 4}).
			WhereNull("deleted_at")

		assert.Equal(
			t,
			"SELECT * FROM users WHERE status = $1 AND user_id IN ($2, $3, $4) AND deleted_at IS NULL",
			b.String(),
		)
		assert.Equal(t, []any{"active", 1, 2, 4}, b.params)
	})
}
