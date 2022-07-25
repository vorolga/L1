package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

func checkUniq(str string) bool {
	m := make(map[rune]int)
	str = strings.ToLower(str)

	arr := []rune(str)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = 1
	}

	return true
}

func main() {
	str := "abcd"

	fmt.Println(checkUniq(str))
}
