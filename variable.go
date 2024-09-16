package main
import("fmt")

// Used multi-line and single line comments and Variable Declaration & Intialization in different ways

/*
func main(){
	var s1 string = "Sathwik"
	var s2 string = "Synectiks"
	x := 2

	fmt.Println("S1:",s1)
	fmt.Println("S2:",s2)
	fmt.Println("x:",x)
}

func main(){
	var a string
	var b int
	var c bool
	var d float32
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main(){
	var s1 string
	s1 = "John"
	fmt.Println(s1)
}

var a  = 2
var b int = 3
var c string = "Sathwik"

func main(){

	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
*/

func main() {

    var a, b, c int = 1, 2, 3
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    var x, y = 1, "Sathwik"
    fmt.Println(x)
    fmt.Println(y)

    e,l := 10,20
    fmt.Println(e)
    fmt.Println(l)

    w, t := 40, "Synectiks" 
    fmt.Println(w)
    fmt.Println(t)

    var (
        g int
        f string = "Sathwik"
        h int = 100
    )
    fmt.Println(g,f,h)
    fmt.Println(f)
    fmt.Println(h)
}