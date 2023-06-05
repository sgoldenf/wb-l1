/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

import "fmt"

type mac struct {
	ports portInterface
}

type portInterface interface {
	connectToUSBCPort()
}

type usbc struct{}

func (uc *usbc) connectToUSBCPort() {
	fmt.Println("Connecting device to USB-C port...")
}

type usb struct{}

func (u *usb) connectToUSBPort() {
	fmt.Println("Connecting device to USB port...")
}

type usbtoUSBCAdapter struct {
	u *usb
}

func (uToC *usbtoUSBCAdapter) connectToUSBCPort() {
	fmt.Println("Connecting USB device to USB-C port via adapter...")
	uToC.u.connectToUSBPort()
}

func main() {
	// Для типа usbc реализованы методы portInterface
	// Значит, можно проинициализировать ports структурой этого типа
	m := &mac{ports: &usbc{}}
	m.ports.connectToUSBCPort()

	// Мы не можем сделать m.ports = &usb{}
	// Для типа usb не реализован метод connectToUSBCPort
	// Можно реализовать структуру-адаптер, которая будет содержать в себе структуру типа usb
	// и при этом удовлетворять интерфейсу portInterface
	m.ports = &usbtoUSBCAdapter{}
	m.ports.connectToUSBCPort()
}
