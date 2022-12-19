package utils

func MapContains[K comparable, V any](theMap map[K]V, key K) bool {
	_, contains := theMap[key]
	return contains
}
