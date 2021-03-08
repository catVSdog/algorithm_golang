package main


func BubbleSort(slice []int){
	for i := len(slice)-1; i > 0; i-- {
		for j := 0; j < i; j ++ {
			if slice[j] > slice[j+1]{
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}