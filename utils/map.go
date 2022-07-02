package utils

type MapCallback[T any, V any] func(T, int) V

func Map[T any, V any](array []T, callback MapCallback[T, V]) []V {
	result := make([]V, 0)
	for i, el := range array {
		thingToAppend := callback(el, i)
		result = append(result, thingToAppend)
	}
	return result
}
