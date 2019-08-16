package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	list := strings.Fields(s)
	for _, word := range list {
		ret[word] += 1
	}
	fmt.Println(ret)
	return ret
}

func main() {
	wc.Test(WordCount)
}

