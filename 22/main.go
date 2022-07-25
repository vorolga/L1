package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

type twoBigInts struct {
	a *big.Int
	b *big.Int
}

func (t *twoBigInts) add() *big.Int {
	return new(big.Int).Add(t.a, t.b)
}

func (t *twoBigInts) mul() *big.Int {
	return new(big.Int).Mul(t.a, t.b)
}

func (t *twoBigInts) sub() *big.Int {
	return new(big.Int).Sub(t.a, t.b)
}

func (t *twoBigInts) div() *big.Int {
	return new(big.Int).Quo(t.a, t.b)
}

func main() {
	ints := &twoBigInts{
		a: new(big.Int),
		b: new(big.Int),
	}
	ints.a, ints.b = ints.a.SetInt64(989326458936993), ints.b.SetInt64(989326458936993)

	fmt.Print("поизведение: ")
	fmt.Println(ints.mul())

	fmt.Print("частное: ")
	fmt.Println(ints.div())

	fmt.Print("сумма: ")
	fmt.Println(ints.add())

	fmt.Print("разность: ")
	fmt.Println(ints.sub())
}
