package main
import ("fmt")

/*
const PI = 3.14

func main(){
	fmt.Println(PI)
}

// Typed constant

const A int = 2
func main(){
	fmt.Println(A)
}

// Untyped Constants

const B = 2

func main(){
	fmt.Println(B)
}

// Unchangeable and Read-Only

func main(){
	const C = 1
	C = 2
	fmt.Println(C)
}
*/

// Multi-Constant Declaration

const(
	a = 3.14
	b = "Sathwik"
	c = 10
)

func main(){
	const(
		e = 20
		f = "Synectiks"
	)

	fmt.Println(a,b,c)
	fmt.Println(e,f)
}


