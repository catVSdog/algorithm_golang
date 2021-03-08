package main

func InsertSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		standardValue := slice[i]
		orderIndex := i - 1
		for orderIndex >= 0 && slice[orderIndex] > standardValue {
			slice[orderIndex+1] = slice[orderIndex]
			orderIndex--
		}
		slice[orderIndex+1] = standardValue
	}
}
