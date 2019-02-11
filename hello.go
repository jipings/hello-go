
package main

import (
	"fmt"
	"sort"
)

func main()  {
	var names []string
	var ages = map[string]int{
		"charlie": 34,
		"alice":   31,
	}
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}