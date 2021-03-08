package main

import (
	"fmt"
)

func main() {
	slice := []int{7, 1, 3, 9, 4, 2, 8, 5, 6}

	// 冒泡排序
	//slice1 := newSlice(slice)
	//BubbleSort(slice1)
	//fmt.Println(slice1)

	// 归并排序
	//slice2 := newSlice(slice)
	//slice2 = MergeSort(slice2)
	//fmt.Println(slice2)

	// 快速排序
	//slice3 := newSlice(slice)
	//QuickSort(slice3, 0 , len(slice3)-1)
	//fmt.Println(slice3)

	// 插入排序
	slice4 := newSlice(slice)
	InsertSort(slice4)
	fmt.Println(slice4)
}

func newSlice(slice []int) []int {
	newSlice := make([]int, 9)
	copy(newSlice, slice)
	return newSlice
}
