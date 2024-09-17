package main
import "fmt"

/*
func main()  {
	
	myslice := [] int {}
	fmt.Println(myslice)
	fmt.Println(cap(myslice))
	fmt.Println(len(myslice))

	myslice2 := [] string {"Sathwik", "Synectiks"}
	fmt.Println(len(myslice2))
  	fmt.Println(cap(myslice2))
  	fmt.Println(myslice2)
}

func main(){

	var my_arr = [5] int {10,20,30,40,50}
	myslice := my_arr[0:3]

	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
}
*/

//using make function
func main()  {
	
	myslice := make([] int, 5, 10)
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice2 := make([] int, 10)
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	myslice = append(myslice, 1, 2, 3, 4, 5, 6)
	fmt.Println(myslice)
	fmt.Println(myslice[3])
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice2 = append(myslice)
	fmt.Println(myslice2)

	myslice3 := make([] int, 11, 20)
	copy(myslice3, myslice2)
	fmt.Println(myslice3)

}