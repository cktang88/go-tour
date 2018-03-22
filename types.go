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

		var s []int = primes[1:4] // can also omit high/low bounds
		s[2] = 999                // note this is actually the 4th elem of the underlying array
		fmt.Println(s)

	}
	{
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
	}
	{
		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		// Slice the slice to give it zero length.
		s = s[:0]
		printSlice(s)

		// extend slice length (see function def below)
		s = extendSliceHandler(s, 4) // should throw (& exit gracefully) if 2nd param > 6
		printSlice(s)

		// Drop its first two values.
		s = s[2:]
		printSlice(s)
	}
	{
		// empty slice = nil
		var s []int
		fmt.Println(s, len(s), cap(s))
		if s == nil {
			fmt.Println("nil!")
		}
	}
	{
		// allocates a zeroed array and returns a slice referring to it
		a := make([]int, 5) // len(a)=5
		// specifying capacity, allows making dynamically-sized arrays
		b := make([]int, 0, 5) // len(b)=0, cap(b)=5

		printSlice(a)
		printSlice(b)
		// NOTE: slices can contain any time, even other slices
	}
	{
		/* appending to slices */
		var s []int
		// append works on nil slices.
		s = append(s, 0)
		// The slice grows as needed.
		s = append(s, 1)
		// We can add more than one element at a time.
		s = append(s, 2, 3, 4)
		printSlice(s)
		// slices always return the new slice

		/* If the backing array of s is too small, a bigger array will be allocated.
		The returned slice will point to the newly allocated array.
		*/
	}

	// for more on slices: https://blog.golang.org/go-slices-usage-and-internals
	{
		// RANGE - returns index & value
		pow := make([]int, 4)
		for i := range pow {
			pow[i] = 1 << uint(i) // == 2**i
		}
		// ignore either var by assigning to _
		for _, value := range pow {
			fmt.Printf("%d\n", value)
			// fmt.Printf(_) // err: cannot use _ as value
		}
	}
	{
		// eg. working with array of arrays
		var ss [][]uint8
		for j := 0; j < dy; j++ {
			var s []uint8
			for i := 0; i < dx; i++ {
				b := i ^ j
				s = append(s, uint8(b))
			}
			ss = append(ss, s)
		}
		return ss
	}
}

/*
Extend its length (assuming sufficient capacity)
If insufficient capacity, will PANIC - then use https://golang.org/doc/effective_go.html#recover
*/
func extendSliceHandler(s []int, newLength int) (newSlice []int) {
	defer func() {
		// recovering from Panic when slice creation fails
		if err := recover(); err != nil {
			// exit gracefully
			fmt.Println("Attempt to extend slice length failed:", err)
		}
	}()
	return s[:newLength]
}

/*
Print slices nicely
*/
func printSlice(s []int) {
	// use len/cap to find length and capacity of slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
