package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter a string:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if strings.HasPrefix(strings.ToLower(input), "i") && strings.Contains(strings.ToLower(input), "a") && strings.HasSuffix(strings.ToLower(input), "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Unfortunately, the string entered either doesn't with 'i', contain 'a', or end with 'n'.")
	}

}
