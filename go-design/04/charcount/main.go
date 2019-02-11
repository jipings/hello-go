// charcount 计算 Unicode 字符个数

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main()  {
	counts := make(map[rune]int) // Unicode 字符数量
	var utflen [utf8.UTFMax + 1]int // UTF-8 的编码长度
	invalid := 0 // 非法UTF-8字符数量

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // 返回 rune、nbytes、error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n",err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n",i,n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}