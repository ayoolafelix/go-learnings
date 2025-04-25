package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the name:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter the address:")
	scanner.Scan()
	address := scanner.Text()

	personLocation := map[string]string{name: address}

	jsonData, err := json.Marshal(personLocation)

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

}
