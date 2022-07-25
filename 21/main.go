package main

import (
	"fmt"
)

//Реализовать паттерн «адаптер» на любом примере.

type phone interface {
	insertUSBTypeC()
}

type IPhone struct {
}

func (i *IPhone) insertUSBTypeC() {
	fmt.Println("Insert USB Type-C into iphone")
}

type Android struct {
}

func (a *Android) insertMicroUSB() {
	fmt.Println("Insert micro USB into android")
}

type AndroidAdapter struct {
	androidPhone *Android
}

func (a *AndroidAdapter) insertUSBTypeC() {
	a.androidPhone.insertMicroUSB()
}

type Client struct {
}

func (c *Client) insertUSBTypeCInPhone(ph phone) {
	ph.insertUSBTypeC()
}
func main() {
	cli := &Client{}
	iphone := &IPhone{}

	cli.insertUSBTypeCInPhone(iphone)

	android := &Android{}

	androidAdapter := &AndroidAdapter{
		androidPhone: android,
	}

	cli.insertUSBTypeCInPhone(androidAdapter)
}
