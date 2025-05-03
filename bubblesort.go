package main

import (
	"fmt"
)

func Swap(arrPtr *[]int, i int) {
	arr := *arrPtr
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func BubbleSort(sliPtr *[]int) {
	sli := *sliPtr
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1-i; j++ {
			if sli[j] > sli[j+1] {
				Swap(sliPtr, j)
			}
		}
	}
}

func main() {

	var numbers []int
	for i := 1; i < 11; i++ {
		var num int
		fmt.Printf("Enter no. %d of 10 integer:", i)
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error reading number:", err)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Before sorting:", numbers)
	BubbleSort(&numbers)
	fmt.Println("After sorting:", numbers)
}
