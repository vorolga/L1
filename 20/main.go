package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

const str = "snow dog sun"

func main() {
	words := strings.Fields(str)

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	for _, item := range words {
		fmt.Print(item, " ")
	}

	fmt.Println()

}
