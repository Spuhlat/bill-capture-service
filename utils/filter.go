package utils

type FilterCallback[T any] func(T, int) bool

func Filter[T any](array []T, callback FilterCallback[T]) []T {
	result := make([]T, 0)
	for i, el := range array {
		if callback(el, i) == true {
			result = append(result, el)
		}
	}
	return result
}
