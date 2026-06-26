package main

import (
	"fmt"

	"store/domain"
)

func main() {
	shop := domain.NewShop()
	shop.Run()
}

//////////////////////////////
func Meaw(){
	
	var meawFunc MeawFunc = func() {
		fmt.Println("meaw")
	}

	meawFunc()
	meawFunc.Meaw()
}

type MeawFunc func()

func (m MeawFunc) Meaw() {
	m()
}

///////////////////////////////////

type IMeaw interface {
	Meaw()
}

type StructMeaw struct{}

func NewStructMeaw() IMeaw {
	return &StructMeaw{}
}

// Meaw implements [IMeaw].
func (s *StructMeaw) Meaw() {
	panic("unimplemented")
}
