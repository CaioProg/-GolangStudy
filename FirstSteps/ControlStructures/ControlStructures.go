package main

import (
	"fmt"
	"os"
)

func main() {

	// If else

	age1 := 18

	if age1 < 13 {
		fmt.Println("You are an child.")
	} else if age1 < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are an adult")
	}

	if age2 := 20; age2 >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	number := 2

	if number%2 == 0 {
		fmt.Println("Is even.")
	} else {
		fmt.Println("Is odd.")
	}

	// For

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	n := 0

	for n < 5 {
		fmt.Println(n)
		n++
	}

	for {
		fmt.Println("This is an infinite loop")
		break // Without this line the loop would continue indefinitely
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fruits := map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}

	for key, value := range fruits {
		fmt.Printf("key: %s, Value: %s\n", key, value)
	}

	str := "hello"

	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	// Switch

	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	age3 := 20

	switch {
	case age3 < 13:
		fmt.Println("You are a child.")
	case age3 < 18:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are an adult.")
	}

	// Defer

	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Panic Recover

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	fmt.Println("Starting the program.")
	panic("Something went wrong!")
}
