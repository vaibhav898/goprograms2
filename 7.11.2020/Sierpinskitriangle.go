package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var order = 4
	var star = "*"

	t := []string{star + strings.Repeat(" ", utf8.RuneCountInString(star))}
	for ; order > 0; order-- {
		sp := strings.Repeat(" ", utf8.RuneCountInString(t[0])/2)
		top := make([]string, len(t))
		for i, s := range t {
			top[i] = sp + s + sp
			t[i] += s
		}
		t = append(top, t...)
	}

	for _, r := range t {
		fmt.Println(r)
	}
}