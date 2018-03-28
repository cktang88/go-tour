/*
interfaces
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
	"io"
)

/*
 * 	INTERFACES - implemented implicitely

 Implicit interfaces decouple the definition of an interface from its implementation,
 which could then appear in any package without prearrangement.

 Some advantages of implicit interfaces:
 https://softwareengineering.stackexchange.com/questions/197356/how-does-go-improve-productivity-with-implicit-interfaces-and-how-does-that-c
*/

/*
 *** Interface Values ***
 Under the covers, interface values can be thought of as a tuple of a value and a concrete type:
 (value, type)
 An interface value holds a value of a specific underlying concrete type.
 Calling a method on an interface value executes the method of the same name on its underlying type.
*/

type I interface {
	M()
}

type T struct {
	S string
}

/*
So type T implements the interface I,
but we don't need to explicitly declare that it does so.

NOTE: If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
WILL NOT lead to a null pointer exception

An interface value that holds a nil concrete value is itself non-nil.
*/
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeInterface(i I) {
	// note templated string
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
A nil interface value holds neither value nor concrete type.
Calling a method on a nil interface is a run-time error
because there is no type inside the interface tuple to indicate which concrete method to call.
*/

/*
Empty interface

Can hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type.
eg, fmt.Print takes any number of arguments of type interface{}.
*/
var i interface{}

func TypeAssertion() {
	var i interface{} = "hello"
	s := i.(string)      // assert: panics if not of type T
	f, ok := i.(float64) // test: "ok" returns whether it is of type T, doesn't panic

	fmt.Println(s, f, ok)

	// type switch (note the test statement)
	switch v := i.(type) {
	case int:
		// here v has type T
	case string:
		// here v has type S
	default:
		// no match; here v has the same type as i
		fmt.Println(v)
	}
}

/*
Stringer is an interface from "fmt"

type Stringer interface {
    String() string
}
*/
/*
Error interface from "fmt"

type error interface {
    Error() string
}
*/
func handleError() {
	i, err := strconv.Atoi("42")
	// most funcs return err, must catch
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

/*
io.Reader interface
Description:
read end of a stream of data.

Go standard library contains many implementations,
including files, network connections, compressors, ciphers, and others.

populates the given byte slice with data
returns the # bytes populated, err val. It returns an io.EOF error when the stream ends.

func (T) Read(b []byte) (n int, err error)
*/

func main() {
	r := strings.NewReader("Hello, Reader!")

	// consumes max 8 bytes at a time
	b := make([]byte, 8)
	// inf loop
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/* By adding Read([]byte) (int, error) to a type, 
you can meka a stream that emits anything */
