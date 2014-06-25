package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "this is really neat"
	words := strings.Split(text, " ")
	for word := range words {
		fmt.Println(words[word])
	}
}
