package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(initialDisplacement, velocity, acceleration float64) func(float64) float64 {
	distance := func(time float64) float64 {
		return initialDisplacement + (velocity * time) + (0.5 * acceleration * math.Pow(time, 2))
	}
	return distance
}

func main() {
	var initialDisplacement, velocity, acceleration, time float64
	fmt.Println("Enter the initial displacement:")
	_, err := fmt.Scan(&initialDisplacement)
	if err != nil {
		fmt.Println("Error reading initial displacement:", err)
		return
	}

	fmt.Println("Enter the velocity:")
	_, err = fmt.Scan(&velocity)
	if err != nil {
		fmt.Println("Error reading velocity:", err)
		return
	}

	fmt.Println("Enter the acceleration:")
	_, err = fmt.Scan(&acceleration)
	if err != nil {
		fmt.Println("Error reading acceleration:", err)
		return
	}

	fmt.Println("Enter the time:")
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Println("Error reading time:", err)
		return
	}

	displacement := GenDisplaceFn(initialDisplacement, velocity, acceleration)
	fmt.Printf("The displacement after %.2f seconds is: %.2f\n", time, displacement(time))
}
