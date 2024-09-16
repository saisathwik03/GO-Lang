package main
import ("fmt")

/*
func main()  {
	
	var i,j string = "Hello", "World!"
	fmt.Print(i,j)
}

func main(){
	var i,j string = "Hello", "World!"
	fmt.Print(i, "\n")
	fmt.Print(j,"\n")
}

func main(){
	var i,j string = "Hello", "World!"
	fmt.Print(i, "\n", j)
}

func main(){
	var i,j string = "Hello", "World!"
	fmt.Print(i, " ", j)
}
*/

func main() {
	var i string = "Hello"
	var j int = 15
  
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
  }