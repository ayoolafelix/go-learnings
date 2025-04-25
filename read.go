package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var fileName string
	fmt.Println("Enter the name of the file to read:")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}

	f, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	everyBody := strings.Split(string(f), "\n")

	type fullName struct {
		fname string
		lname string
	}

	people := make([]fullName, 0)
	var someBody fullName

	for _, person := range everyBody {

		if strings.TrimSpace(person) == "" {
			continue
		}
		someBody.fname = strings.Split(person, " ")[0]
		someBody.lname = strings.Split(person, " ")[1]
		people = append(people, someBody)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("First Name:", people[i].fname)
		fmt.Println("Last Name:", people[i].lname, "\n")
	}

}
