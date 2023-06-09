package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var str string
	var substr string
	flag.StringVar(&str, "str", "строка для поиска", "string to search in")
	flag.StringVar(&substr, "substr", "строка", "string to search for")
	flag.Parse()
	fmt.Println(strings.Contains(str, substr))
}
