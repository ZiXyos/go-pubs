package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SliceUpscale[K comparable](array *[]K) {
  currLen := len(*array);
  currCap := cap(*array)
	if currLen >= (currCap / 2) {
		newSlice := make([]K, currLen, (currCap * 2 + currLen));
		copy(newSlice, *array);
		*array = newSlice;
	}
}

func SliceDownscale[K comparable](arr *[]K) {
  currLen := len(*arr);
  currCap := cap(*arr)
  if currLen < ((currCap + 10) / 2) {
    newSlice := make([]K, currLen, (currCap / 2 + currLen));
    copy(newSlice, *arr);
    *arr = newSlice;
  }
}

func SortedInsert[K constraints.Ordered](arr *[]K, v K) {
  idx := sort.Search(len(*arr), func(i int) bool { return (*arr)[i] >= v })
  *arr = append((*arr)[:idx], append([]K{v}, (*arr)[idx:]...)...)
}

