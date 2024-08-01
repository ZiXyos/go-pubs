package utils

func SliceUpscale[K comparable](array *[]K) {
	if len(*array) >= (cap(*array) / 2) {
		newSlice := make([]K, len(*array), (cap(*array) * 2 + len(*array)));
		copy(newSlice, *array);
		*array = newSlice;
	}
}

func SliceDownscale[K comparable](arr *[]K) {
  if len(*arr) < ((cap(*arr) + 10) / 2) {
    newSlice := make([]K, len(*arr), (cap(*arr) / 2 + len(*arr)));
    copy(newSlice, *arr);
    *arr = newSlice;
  }
}

