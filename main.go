package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {

	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	//fmt.Println(multiply(2, 3))
	//totalLength, upperName := lenAndUpper("sungbeen")
	//fmt.Println(totalLength)
	//fmt.Println(upperName)

	repeatMe("a", "b", "dd", "d", "e", "f")
}
