package main
import "fmt"

func familyName(fname string) {
	fmt.Println("Hello", fname, "Welcome to Synectiks !!")
}

func addition_of_two_numbers(x int, y int) int {
	return x+y
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
  
func main() {
	familyName("Sathwik")
	fmt.Println(addition_of_two_numbers(1,2))
	a, _ := myFunction(5, "Hello")
  	fmt.Println(a)
}