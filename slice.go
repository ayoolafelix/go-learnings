package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make([]int, 3)
	var newNumber string
	for {
		fmt.Printf("Enter a number to begin or 'X' to exit the program:")
		_, err := fmt.Scanln(&newNumber)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if strings.ToUpper(newNumber) == "X" {
			fmt.Println("Exiting the program.")
			break
		} else {
			num, err := strconv.Atoi(newNumber)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				continue
			}
			numbers = append(numbers, num)
			sort.Ints(numbers)
			fmt.Println(numbers)
		}
	}

}
