package util

import "math/rand"

func UniqueIntSlice(slice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func AllFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func SomeFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func RandomIntSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
