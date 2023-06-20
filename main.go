package main

import "fmt"

func main() {
	foods := []string{"ramen", "kimchi", "sushi"}
	fmt.Printf("%v\n", foods)
	foods = append(foods, "rice")
	fmt.Print(foods)
}
