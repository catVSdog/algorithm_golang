package main

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	leftArray := MergeSort(slice[:mid])
	rightArray := MergeSort(slice[mid:])
	return merge(leftArray, rightArray)
}

func merge(left, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	mergedSlice := []int{}

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			mergedSlice = append(mergedSlice, left[leftIndex])
			leftIndex++
		} else {
			mergedSlice = append(mergedSlice, right[rightIndex])
			rightIndex++
		}
	}
	mergedSlice = append(mergedSlice, left[leftIndex:]...)
	mergedSlice = append(mergedSlice, right[rightIndex:]...)
	return mergedSlice
}
