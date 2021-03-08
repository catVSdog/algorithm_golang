package main

func SelectSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		if min != i {
			slice[i], slice[min] = slice[min], slice[i]
		}
	}
}
