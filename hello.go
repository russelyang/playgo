package main

import (
	"t2gp.com/oauth2/login/morestrings"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
