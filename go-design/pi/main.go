
package main

import (
	"fmt"
	"math"
	"time"
)

func main()  {
	start := time.Now()
	fmt.Printf("%f \n",pi(500000))
	fmt.Println(time.Since(start))
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k<= n; k++ {
		go term(ch, float64(k))
	}
	f :=0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4*math.Pow(-1, k)/(2*k+1)
}