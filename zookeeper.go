package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type snake struct {
	meal  string
	move  string
	sound string
}

func (s snake) Eat() string {
	return s.meal
}

func (s snake) Move() string {
	return s.move
}

func (s snake) Speak() string {
	return s.sound
}

type cow struct {
	meal  string
	move  string
	sound string
}

func (c cow) Eat() string {
	return c.meal
}

func (c cow) Move() string {
	return c.move
}

func (c cow) Speak() string {
	return c.sound
}

type bird struct {
	meal  string
	move  string
	sound string
}

func (b bird) Eat() string {
	return b.meal
}

func (b bird) Move() string {
	return b.move
}

func (b bird) Speak() string {
	return b.sound
}

func main() {
	zoo := map[string]string{}

	for {
		fmt.Println("Enter a query type, a name, and a request (e.g., 'newanimal cow bird'):")

		reader := bufio.NewReader(os.Stdin)
		animalQuery, _ := reader.ReadString('\n')
		animalQuery = strings.TrimSpace(animalQuery)

		userInput := strings.Fields(animalQuery)

		if len(userInput) < 3 {
			fmt.Println("Invalid input. Please provide a query type, a name, and a request.")
			continue
		}

		queryType := userInput[0]
		animalName := userInput[1]
		userQuery := userInput[2]

		if queryType == "newanimal" {
			zoo[animalName] = userQuery
			fmt.Printf("Created a new animal: %s\n", animalName)
		} else if queryType == "query" {
			animalType, ok := zoo[animalName]
			if !ok {
				fmt.Printf("Animal %s not found.\n", animalName)
				continue
			}

			var animal Animal
			switch animalType {
			case "cow":
				animal = cow{meal: "Grass", move: "Walk", sound: "Moo"}
			case "bird":
				animal = bird{meal: "Worms", move: "Fly", sound: "Peep"}
			case "snake":
				animal = snake{meal: "Mice", move: "Slither", sound: "Hsss"}
			default:
				fmt.Println("Unknown animal type.")
				continue
			}

			actions := map[string]func(Animal) string{
				"eat":   Animal.Eat,
				"move":  Animal.Move,
				"speak": Animal.Speak,
			}

			if action, ok := actions[userQuery]; ok {
				fmt.Println(action(animal))
			} else {
				fmt.Println("Unknown query. Please use 'eat', 'move', or 'speak'.")
			}
		} else {
			fmt.Println("Unknown query type. Please use 'newanimal' or 'query'.")
		}
	}
}
