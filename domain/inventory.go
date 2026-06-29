package domain

import (
	"fmt"
)

type inventory struct {
	Money    int `json:"money"`
	Products map[string]*[]Product
}

func (i *inventory) Show() {
	number := 1

	fmt.Println("Money: ", i.Money)

	for _, products := range i.Products {
		for _, product := range *products {
			fmt.Println("Shopping list: ", number, product.Name, product.Price)
			number++
		}
	}
}
