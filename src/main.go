package main // must have main for entry

import ( // unused pkg cause error
	"fmt"
	"os"
)

func getNumber() int {
	return 8282
}

func log(message string) { // void

}

// func add(a,b int) int // short
func add(a int, b int) int { // 2 in 1 out
	return 0 // 1 out
}

func power(name string) (int, bool) { // 1 in 2 out
	return 0, false // 2 out
}

func main() { // must called main for entry point
	// Args[0] is path to current bin
	if len(os.Args) > 2{ // size of str, dict, array
		os.Exit(1)
	}

	var number int // declare
	number = 82 // assign
	var num string = "num" // unused var cause error
	n, nfunc := 2, getNumber() // multi-assign
	num, newvar := "samenum", "newvar"// at least one new var to use :=

	if len(os.Args) == 2{
		fmt.Println("Args[1]:", os.Args[1])
	}

	fmt.Printf("number:%d\n", number) // print format
	fmt.Printf("num:%s n:%d \n", num, n)
	fmt.Printf("nfunc:%d newvar:%s\n", nfunc, newvar)

	_, exists := power("wh") // _ not assign, any type
	if exists == false {
		// error
	}
}