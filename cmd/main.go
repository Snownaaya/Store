package main

import "store/domain"

func main() {
	shop := domain.NewShop()
	shop.Run()
}