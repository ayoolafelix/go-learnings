package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name  string
	meal  string
	move  string
	sound string
}

func (ani Animal) Eat() string {
	return fmt.Sprintf("%s eats %s.", ani.name, ani.meal)
}

func (ani Animal) Move() string {
	return fmt.Sprintf("%s moves by %s.", ani.name, ani.move)
}

func (ani Animal) Speak() string {
	return fmt.Sprintf("%s says %s.", ani.name, ani.sound)
}

func main() {
	cow := Animal{name: "Cow", meal: "Grass", move: "Walk", sound: "Moo"}
	bird := Animal{name: "Bird", meal: "Worms", move: "Fly", sound: "Peep"}
	snake := Animal{name: "Snake", meal: "Mice", move: "Slither", sound: "Hsss"}

	for {
		fmt.Println("Enter an animal and a query (e.g., 'cow eat'):")

		reader := bufio.NewReader(os.Stdin)
		animalQuery, _ := reader.ReadString('\n')
		animalQuery = strings.TrimSpace(animalQuery)

		userInput := strings.Fields(animalQuery)

		if len(userInput) < 2 {
			fmt.Println("Invalid input. Please provide an animal and a query (e.g., 'cow eat').")
			return
		}

		animalName := userInput[0]
		useQuery := userInput[1]

		actions := map[string]func(Animal) string{
			"eat":   Animal.Eat,
			"move":  Animal.Move,
			"speak": Animal.Speak,
		}

		if action, ok := actions[useQuery]; ok {
			switch animalName {
			case "cow":
				fmt.Println(action(cow))
			case "bird":
				fmt.Println(action(bird))
			case "snake":
				fmt.Println(action(snake))
			default:
				fmt.Println("Unknown animal")
			}
		} else {
			fmt.Println("Unknown query. Please use 'eat', 'move', or 'speak'.")
		}
	}
}
