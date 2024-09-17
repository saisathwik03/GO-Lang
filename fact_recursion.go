package main
import "fmt"

func fact(x float64) (y float64) {
	if x > 0 {
		y = x * fact(x - 1)
	} else {
		y = 1
	}
	return
}

func main() {
	var x = float64(5) 
	fmt.Println(fact(x))
}
