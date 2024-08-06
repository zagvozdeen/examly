package util

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
