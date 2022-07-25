package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]int)
	for _, v := range arr {
		m[v] = 1
	}
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
}
