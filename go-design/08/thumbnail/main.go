// Run with:
//   $ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//   foo.jpeg
//   ^D
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string)  {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func main()  {
	input := bufio.NewScanner(os.Stdin)
	 var filenames []string;
	for input.Scan() {
		// thumb, err := thumbnail.ImageFile(input.Text())
		// if err != nil {
		// 	log.Print(err)
		// 	continue
		// }
		// fmt.Println(thumb)
		filenames = append(filenames, input.Text())
		fmt.Println(filenames)
		makeThumbnails(filenames)
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}