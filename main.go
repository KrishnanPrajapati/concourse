package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on Slices")

	var fruitLess = []string{"Apple", "Tomato"}
	fruitLess = append(fruitLess, "Mango", "Banana")

	fruitLess = append(fruitLess[1:3]) // It will Get Rane between data
	highscores := make([]int, 4)

	highscores[0] = 224
	highscores[1] = 243
	highscores[2] = 274
	highscores[3] = 254

	highscores = append(highscores, 748, 875, 784, 343, 434)
	fmt.Println(highscores)

	sort.Ints(highscores)
	//How to remove a value from slices based on index
	var courses = []string{"Ruby", "Pahthon", "readtJS", "Javascript", "Swift"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)
}
