package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse1@gmail.com"

func main()  {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9].+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)
}

// FindAllString 获取所有的匹配
// FindAllStringSubmatch 匹配子串