package util

import (
	"github.com/Den4ik117/examly/internal/model"
)

func GetMapByIDFromSlice[T model.IModel](models []T) map[int]T {
	r := make(map[int]T)
	for _, m := range models {
		r[m.GetID()] = m
	}
	return r
}
