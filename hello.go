package main

// import "unsafe";
import "fmt";

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)
func main()  {
	fmt.Println(i,j,k,l)
}