package main
import "fmt"

func main()  {
	
	var cars = [3] string {"BMW", "AUDI", "XUV700"}
	arr1 := [5] int {10,20,30}
	arr2 := [5] int {10,20,30,40,50}
	var arr3 = [3] int {}
	arr4 := [3]int {1:10, 2:20}

	fmt.Println(cars)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(len(cars))
	fmt.Println(arr1[1])
}