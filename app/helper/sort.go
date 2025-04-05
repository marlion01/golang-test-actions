package helper

import "sort"

func Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func SortString(arr []string) []string {
	sort.Strings(arr)
	return arr
}

func SortFloat(arr []float64) []float64 {
	sort.Float64s(arr)
	return arr
}
