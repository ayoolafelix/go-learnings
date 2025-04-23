package main

import "fmt"

func main() {
	var toTruncate float64

	fmt.Printf("Enter a decimal number: ")
	_, err := fmt.Scan(&toTruncate)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Truncated number: %d\n", int(toTruncate))
}
