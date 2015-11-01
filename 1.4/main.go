package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	str = "aabc"
	fmt.Printf("input %s\noutput %s\n", str, compress(str))
	str = "abc"
	fmt.Printf("input %s\noutput %s\n", str, compress(str))
	str = "abccccc"
	fmt.Printf("input %s\noutput %s\n", str, compress(str))
}

func compress(a string) string {
	chars := strings.Split(a, "")
	prev := ""
	count := 0
	result := ""
	for _, ch := range chars {
		if prev != "" && prev != ch {
			if count > 1 {
				result += fmt.Sprintf("%s%d", prev, count)
			} else {
				result += prev
			}
			count = 1
		} else {
			count++
		}
		prev = ch
		//		fmt.Printf("c:%s p:%s r:%s cnt:%d\n", ch, prev, result, count)
	}
	if count > 1 {
		result += fmt.Sprintf("%s%d", prev, count)
	} else {
		result += prev
	}
	return result
}
