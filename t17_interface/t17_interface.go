package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone is working")
}

func (p Phone) Stop() {
	fmt.Println("Phone is not working")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera is working")
}

func (c Camera) Stop() {
	fmt.Println("Camera is not working")
}

type Computer struct {
}

func (c Computer) working(u Usb) {
	//通过usb接口来调用两个方法
	u.Start()
	u.Stop()
}

func main() {
	phone := Phone{}
	computer := Computer{}
	camera := Camera{}

	computer.working(phone)
	computer.working(camera)

}
