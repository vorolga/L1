package main

import (
	"fmt"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

//Как я поняла задание, надо инвертировать i бит
func main() {
	var data int64
	var i uint64

	fmt.Println("Введите число и номер бита")
	_, _ = fmt.Scan(&data, &i)

	data = data ^ 1<<i
	fmt.Println(data)
}
