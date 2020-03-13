package main

import "fmt"

type animal interface {
	Run()
}

type dog struct {
}

func newDog() animal {
	return &dog{}
}

func (*dog) Run() {
	fmt.Println("wan")
}

type cat struct {
}

func newCat() animal {
	return &cat{}
}

func (*cat) Run() {
	fmt.Println("niao")
}

func interf() {
	newDog().Run()
	newCat().Run()
}
