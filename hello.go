
package main

import (
	"fmt"
)

func f() func(int) int {
	i := 0
    return func(r int) int {
		s := i+r
		return s
	}
}

func main()  {
	s := f()
	defer fmt.Println(s(1))
	defer fmt.Println(s(2))
	defer fmt.Println(s(3))
	fmt.Println(s(5))
	
}