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
	slice2 := newSlice(slice)
	slice2 = MergeSort(slice2)
	fmt.Println(slice2)

}

func newSlice(slice []int)[]int{
	newSlice := make([]int, 9)
	copy(newSlice, slice)
	return newSlice
}