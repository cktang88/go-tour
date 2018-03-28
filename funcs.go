/*
functions
*/

package main

import (
	"fmt"
	"math"
)

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

/*
Supports lambdas (called function literals, https://golang.org/ref/spec#Function_literals)
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func compute_test() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

/*
Closures
A closure is a function value that references variables from outside its body.
So, it is "bound" to those variables
Allows each function to have it's own state, like a mini-class
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int { // this func literal is a closure
		sum += x
		return sum
	}
}

func adder_test() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/* Closure Excercise:
 * Returns a closure
 */
func fibonacci() func() int {
	a, b, n := 0, 1, 1
	return func() int {
		switch n % 2 {
		case 0:
			a = a + b
		case 1:
			b = b + a
		}
		n = n + 1
		if n%2 == 0 {
			return a
		}
		return b
	}
}

func fib_test() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/* NO CLASSES. But can define METHODS on types.
A method is a function with a special RECEIVER argument.
*/
type Vertex struct {
	X, Y float64
}

/* RECEIVER (eg. Vertex) is right before func name (eg. Abs)

func sig similar to regular "func Abs(v Vertex) float64 {", except more convenient usage
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func abs_test() {
	v := Vertex{3, 4}
	// so now method can kind of be used like method of a class
	fmt.Println(v.Abs())
}

/* NOTE: Can also define methods on non-struct types
eg.
type MyFloat float64
func (f MyFloat) Abs() float64 {...}

Usage: f.Abs()

*** Only works if receiver type is declared in same package as method
*/

/* Pointer receivers

Can be called with regular Vertex v, eg. (v.Scale) gets auto converted to (&v).Scale
*/
func (v *Vertex) Scale(f float64) {
	// can modify receiver
}

/* Therefore, cannot declare below, b/c would be interpreted as SAME method as above ptr method

func (v Vertex) Scale(f float64) {
}
*/
func ptr_method_test() {
	v := Vertex{3, 4}
	v.Scale(.5)
}

/*
NOTE:
All methods on a given type should have EITHER value or pointer receivers, but not a mixture of both
*/
