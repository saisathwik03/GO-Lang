package main

import "fmt"

// Function to check map values and existence of keys
func check() {
	var a = map[string]string{
		"brand": "Ford",
		"model": "Mustang",
		"year":  "1964",
		"day":   "",
	}

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}

// Function to demonstrate map aliasing
func reference() {
	var a = map[string]string{
		"brand": "Ford",
		"model": "Mustang",
		"year":  "1964",
	}
	b := a // b is an alias for a

	fmt.Println("Original maps:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b["year"] = "1970" // Modifying b will also modify a

	fmt.Println("After change to b:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

// Function to iterate over a map
func iterate() {
	a := make(map[string] int)
	a["one"] = 1
	a["two"] = 2

	for k, v := range a {
		fmt.Printf("%v: %v\n", k, v)
	}
}

// Function to iterate over a map with a predefined list of keys
func iterate_order() {
	a := map[string]int{
		"one": 1,
		"two": 2,
	}
	var b []string

	b = append(b, "one", "two")

	for k, v := range a {
		fmt.Printf("%v: %v", k, v)
	}

	fmt.Println()

	for _, ele := range b {
		fmt.Printf("%v: %v", ele, a[ele])
	}
}

func main() {
	check()
	reference() // aliasing
	iterate()
	iterate_order()
}