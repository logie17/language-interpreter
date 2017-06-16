package parser

import (
	"fmt"
	"strings"
)

var tabs = 0
var enabled = false
func trace(o string) string {
	if (enabled) {
		fmt.Printf("%sBEGIN %s\n", strings.Repeat("\t", tabs), o)
		tabs++
	}
	return o
}

func untrace(o string) {
	if (enabled) {
		tabs--
		fmt.Printf("%sEND %s\n", strings.Repeat("\t", tabs), o)
	}
}
