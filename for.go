package main
import "fmt"

/*
func main()  {
	
	adj := [] string {"small", "sweety"}
	fruits := [] string {"apple", "orange", "banana"}

	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}
*/

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for i, j := range fruits {
     fmt.Printf("%v\t%v\n", i, j)
  }
}