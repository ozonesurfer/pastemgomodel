// languages
package model

import (
//	"fmt"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/
type Language struct {
	Id   int64
	Name string
	//	Keywords []string
}

var Languages = []Language{
	Language{int64(0), "Text"},
	Language{int64(1), "Go"},
	Language{int64(2), "JavaScript"},
	Language{int64(3), "Java"},
	Language{int64(4), "PHP"},
	Language{int64(5), "Ruby"},
	Language{int64(6), "C#"},
}
