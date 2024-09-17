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

}