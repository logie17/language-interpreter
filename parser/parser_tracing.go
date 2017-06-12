package parser

import (
	"fmt"
	"strings"
)

var tabs = 0

func trace(o string) string {
	fmt.Printf("%sBEGIN %s\n", strings.Repeat("\t", tabs), o)
	tabs++
	return o
}

func untrace(o string) {
	tabs--
	fmt.Printf("%sEND %s\n", strings.Repeat("\t", tabs), o)
}
