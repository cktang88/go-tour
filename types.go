/*
Types: structs, arrays, slices, and maps.
*/

package main

import "fmt"

func main() {
	// POINTERS (C-like syntax), but NO POINTER ARITHMETIC
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


	// structs
	type Vertex struct {
		X int
		Y int
	}
	// for nested structs: https://stackoverflow.com/questions/24809235/initialize-a-nested-struct-in-golang
	{
		v := Vertex{1, 2}
		// dot notation
		v.X = 4
		fmt.Println(v.X)
		// can also use pointer to struct, but can still use dot notation without explicit dereference
		p := &v
		p.X = 1e6
		fmt.Println(v.X)
	}
	{
		//	shorthand
		v1 := Vertex{1, 2}  // has type Vertex
		v2 := Vertex{X: 1}  // Y:0 is implicit
		v3 := Vertex{}      // X:0 and Y:0
		v4 := &Vertex{1, 2} // has type *Vertex
		fmt.Println(v1, v2, v3, v4)
	}
	{
		// arrays
		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
		fmt.Println(a[0], a[1])
		fmt.Println(a)
	
		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)

		/* Slicing
		* A slice doesn't store data, just stores reference to array.
		* Changing elems in slice changes underlying array
		*/

		var s []int = primes[1:4]
		s[2] = 999 // note this is actually the 4th elem of the underlying array
		fmt.Println(s)

		// Slice literals - is another way to create a slice. Like an array literal, without the length. 
		r := []bool{true, false, true, true, false, true}
		fmt.Println(r)
		// note initializing array of structs
		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)
		// for more on slices: https://blog.golang.org/go-slices-usage-and-internals
	}

}
