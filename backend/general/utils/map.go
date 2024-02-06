package utils

import (
	"sort"
)

func SortByValue(m map[string]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		num1 := StringToInt64(m[keys[i]])
		num2 := StringToInt64(m[keys[j]])
		return num1 > num2
	})
	return keys
}

func SortStringSlice(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		num1 := StringToInt64(list[i])
		num2 := StringToInt64(list[j])
		return num1 > num2
	})
	return list
}

func SortByIntValue(m map[string]int64) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}
