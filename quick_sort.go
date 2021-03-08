package main

func QuickSort(slice []int, begin, end int) {
	if begin > end {
		return
	}
	leftCur := begin
	rightCur := end
	standardValue := slice[end]

	for leftCur < rightCur {
		for leftCur < rightCur && slice[leftCur] < standardValue {
			leftCur++
		}
		slice[rightCur] = slice[leftCur]
		for leftCur < rightCur && slice[rightCur] > standardValue {
			rightCur--
		}
		slice[leftCur] = slice[rightCur]
	}

	slice[leftCur] = standardValue
	QuickSort(slice, begin, leftCur-1)
	QuickSort(slice, leftCur+1, end)
}
