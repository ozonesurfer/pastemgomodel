// languages
package model

import (
	//	"fmt"
	"sort"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/
type Language struct {
	Id       int64
	Name     string
	Keywords sort.StringSlice
}

var Languages = []Language{
	Language{int64(0), "Text", nil},
	Language{int64(1), "Go",
		sort.StringSlice{
			"break", "case", "chan", "const",
			"continue", "default", "defer", "else", "fallthrough", "for", "func",
			"go", "goto", "if", "import", "interface", "map", "package", "range",
			"return", "select", "struct", "switch", "type", "var",
		},
	},
}
