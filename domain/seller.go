package domain

import (
	"fmt"
	"strconv"

	"store/extensionflow"
)

type Seller struct {
	inventory
}

func NewSeller() *Seller {
	products := make(map[string]*[]Product)

	products["Electronics"] = &[]Product{
		{Name: "Phone", Price: 1000},
		{Name: "Monitor", Price: 2000},
		{Name: "Keyboards", Price: 500},
		{Name: "Laptop", Price: 5000},
		{Name: "Mouse", Price: 300},
		{Name: "WashingMachine", Price: 10000},
		{Name: "Fan", Price: 900},
		{Name: "Iron", Price: 200},
		{Name: "DigitalCameras ", Price: 1500},
	}

	return &Seller{
		inventory: inventory{
			Products: products,
			Money:    1000,
		},
	}
}

func (s *Seller) Sell(index int) {
	products := s.Products["Electronics"]

	product := (*products)[index]
	s.Money += product.Price

	*products = append(
		(*products)[:index],
		(*products)[index+1:]...)

	s.Show()
}

func (s *Seller) Input() int {
	input := extensionflow.UserInput("\nEnter the name of the potion you buy: ")
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: You entered a non-numeric value. Please try again.")
	}

	index := number - 1
	return index
}

func (s *Seller) HasProduct(name string) bool {
	for _, categoryItems := range s.Products {
		for _, item := range *categoryItems {
			if item.Name == name {
				return true
			}
		}
	}
	return false
}

func (s *Seller) TakeMoney(product Product) {
	s.Money += product.Price
}
