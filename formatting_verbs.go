package main
import "fmt"

/*

// General Formatting
func main()  {
	
	var i = 10.5
	var j string = "Hello"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", j)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%T\n", j)
}

// Integer formatting
func main() {
	var i = 15
   
	fmt.Printf("%b\n", i)	// base 2
	fmt.Printf("%d\n", i)	// base 10
	fmt.Printf("%+d\n", i)	// base 10 always show sign
	fmt.Printf("%o\n", i)	// base 8
	fmt.Printf("%O\n", i)	// base 8 leding with Oo
	fmt.Printf("%x\n", i)	// base 16 lower case
	fmt.Printf("%X\n", i)	// base 16 upper case
	fmt.Printf("%#x\n", i)	// base 16 leading with Ox
	fmt.Printf("%4d\n", i)	// pad with 4 spces width right
	fmt.Printf("%-4d\n", i)	// pad with 4 spaces width left
	fmt.Printf("%04d\n", i)	// pad with zeros width 4
  }


// Sring fomratting
func main() {
	var j = "Hello"
   
	fmt.Printf("%s\n", j)	// plain string
	fmt.Printf("%q\n", j)	// double quoted string
	fmt.Printf("%8s\n", j)	// plain string with pad 8 right
	fmt.Printf("%-8s\n", j)	// plain string with pad 8 left
	fmt.Printf("%x\n", j)	// value in hex format
	fmt.Printf("% x\n", j)	// va;ue in hex format with spaces
  }
*/

func main() {
	var i = 3.141
  
	fmt.Printf("%e\n", i)	//exponent value
	fmt.Printf("%f\n", i)	// no exponent value
	fmt.Printf("%.2f\n", i)	//precision 2 values
	fmt.Printf("%6.2f\n", i)	// 6 width 2 precision
	fmt.Printf("%g\n", i)	// only necessary digits
  }
  