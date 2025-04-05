package helper_test

import (
	"sort"
	"testing"

	"example.com/helper"
)

func TestSort(t *testing.T) {
	arr := []int{3, 2, 1}
	arr = helper.Sort(arr)
	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Errorf("Sort failed")
	}
}

func TestSortString(t *testing.T) {
	arr := []string{"c", "b", "a"}
	sort.Strings(arr)
	if arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Errorf("SortString failed")
	}
}
