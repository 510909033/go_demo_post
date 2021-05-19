package main

func maopao1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < (len(arr) - i); j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
func maopao2(arrays []int) {
	for i := 0; i < len(arrays); i++ {
		for j := 1; j < len(arrays)-i; j++ {
			if arrays[j-1] > arrays[j] {
				arrays[j], arrays[j-1] = arrays[j-1], arrays[j]
			}
		}
	}
}
