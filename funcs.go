/*
functions
*/

package main

import "fmt"

// multiple named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	/* long assigment, is the only type that can be used
	outside funcs bc only reserved keywords (eg var) can be outside funcs */
	var _x, _y int = split(17)
	// short assignment, uses type inference, can only be used inside functions
	k, v := 5, true
	const IsThisTrue = true // const, can't be declared via :=
	fmt.Println(IsThisTrue)
	fmt.Println(k + 1)
	fmt.Println(_x, _y, v)
}
