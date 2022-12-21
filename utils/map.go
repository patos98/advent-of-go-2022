package utils

func MapContains[K comparable, V any](theMap map[K]V, key K) bool {
	_, contains := theMap[key]
	return contains
}

func CopyMap[K comparable, V any](mapToCopy map[K]V, targetMap map[K]V) map[K]V {
	for k, v := range mapToCopy {
		targetMap[k] = v
	}
	return targetMap
}
