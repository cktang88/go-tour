/*
Control flow: If, Else, For, Switch, Defer
*/

package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

func main() {
	sum := 0
	if bounds := 0; sum >= bounds { // if loop, note: no parens, can put initial statement beforehand, like for loop
		// regular for-loop
		for i := 0; i < 10; i++ {
			sum += i
		}
	} else { // having this on separate line is actually syntax error!
		// else
		fmt.Println("hello")
	}
	fmt.Println(sum)

	// stuff
	{
		// pre and post optional
		sum := 1
		for sum < 1000 {
			sum += sum
		}
	}
	{
		// "while" loop
		sum := 1
		for sum < 1000 {
			sum += sum
		}
	}
	{
		// inf loop
		/*
			for {
			}
		*/
	}
	{
		// switch
		// eg. shows OS being used
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.", os)
		}
	}
	{
		// switch without var = substitute for if/else chains
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
		}
	}
	/* DEFER - defers statement execution until surrounding function returns.
	A defer statement pushes a function call onto a list.
	The list of saved calls is executed after the surrounding function returns.
	*/

	defer fmt.Println("Program ended. :)")
	fmt.Println("just before prog ends...")
}

func CopyFile(dstName, srcName string) (written int64, err error) {

	/* example of stacking defers
	Deferred function calls are pushed onto a stack.
	When a function returns, its deferred calls are executed in last-in-first-out order.

	// enables some (weak?) EXCEPTION SAFETY --> files always closed if error occurs b/c func will exit
	*/
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
