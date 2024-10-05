package util

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	b        *strings.Builder
	parts    []string
	params   []any
	children bool
}

func NewQueryBuilder(s string) *QueryBuilder {
	b := &strings.Builder{}
	b.WriteString(s)
	return &QueryBuilder{
		b: b,
	}
}

func (b *QueryBuilder) WhereFunc(fn func(*QueryBuilder)) *QueryBuilder {
	sb := &QueryBuilder{
		b:        &strings.Builder{},
		params:   b.params,
		children: true,
	}
	fn(sb)
	if len(b.parts) != 0 {
		b.parts = append(b.parts, "AND")
	}
	b.parts = append(b.parts, fmt.Sprintf("(%s)", sb.String()))
	b.params = sb.params
	return b
}

func (b *QueryBuilder) OrWhereFunc(fn func(*QueryBuilder)) *QueryBuilder {
	sb := &QueryBuilder{
		b:        &strings.Builder{},
		params:   b.params,
		children: true,
	}
	fn(sb)
	b.parts = append(b.parts, "OR")
	b.parts = append(b.parts, fmt.Sprintf("(%s)", sb.String()))
	b.params = sb.params
	return b
}

func (b *QueryBuilder) WhereNull(column string) *QueryBuilder {
	if len(b.parts) != 0 {
		b.parts = append(b.parts, "AND")
	}
	b.parts = append(b.parts, fmt.Sprintf("%s IS NULL", column))
	return b
}

func (b *QueryBuilder) Where(column string, comparable string, value any) *QueryBuilder {
	start := len(b.params) + 1
	if len(b.parts) != 0 {
		b.parts = append(b.parts, "AND")
	}
	b.parts = append(b.parts, fmt.Sprintf("%s %s $%d", column, comparable, start))
	b.params = append(b.params, value)
	return b
}

func (b *QueryBuilder) OrWhere(column string, comparable string, value any) *QueryBuilder {
	start := len(b.params) + 1
	b.parts = append(b.parts, "OR")
	b.parts = append(b.parts, fmt.Sprintf("%s %s $%d", column, comparable, start))
	b.params = append(b.params, value)
	return b
}

func (b *QueryBuilder) WhereIn(column string, values []any) *QueryBuilder {
	start := len(b.params) + 1
	if len(b.parts) != 0 {
		b.parts = append(b.parts, "AND")
	}
	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+start)
	}
	b.parts = append(b.parts, fmt.Sprintf("%s IN (%s)", column, strings.Join(placeholders, ", ")))
	b.params = append(b.params, values...)
	return b
}

func (b *QueryBuilder) String() string {
	if !b.children {
		b.b.WriteString(" WHERE ")
	}
	b.b.WriteString(strings.Join(b.parts, " "))
	return b.b.String()
}

func (b *QueryBuilder) Params() []any {
	return b.params
}
