package utils

import (
	"sort"
)

func InsertIntoSortedSlice[T any](items []T, itemToInsert T, emptyItem T, comparator func(item T, itemToInsert T) bool) []T {
	i := sort.Search(len(items), func(i int) bool { return comparator(items[i], itemToInsert) })
	items = append(items, emptyItem)
	copy(items[i+1:], items[i:])
	items[i] = itemToInsert
	return items
}

func SliceContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
