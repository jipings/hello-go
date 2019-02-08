
// 找出重复的行
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main()  {
	counts := make(map[string]int);
	fmt.Println(os.Args[1:])
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		fmt.Println(strings.Split(string(data), "\n"))
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n >1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// go run 1.3.go ./data.text 