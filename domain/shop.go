package domain

import (
	"fmt"
	"strconv"

	"store/extensionflow"
)

type Shop struct {
	customer *Customer
	seller   *Seller
	commands map[string]func()
	isWork   bool
}

const (
	ShowInfo         string = "1"
	SellCommand      string = "2"
	InventoryCommand string = "3"
	ExitCommand      string = "4"
)

func NewShop() *Shop {
	shop := &Shop{
		customer: newCustomer(),
		seller:   NewSeller(),
	}

	fmt.Println("1 - ShowPotions")
	fmt.Println("2 - Sell")
	fmt.Println("3 - Inventory")
	fmt.Println("4 - Exit")

	return shop
}

func (s *Shop) Run() {

	for s.isWork == true {
		input := extensionflow.UserInput("Enter command: ")

		s.commands = map[string]func(){
			ShowInfo:         func() { s.seller.Show() },
			SellCommand:      func() { s.sellProductFlow() },
			InventoryCommand: func() { s.customer.Show() },
			ExitCommand:      func() { s.isWork = false },
		}

		if cmd, ok := s.commands[input]; ok {
			cmd()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func (s *Shop) sellProductFlow() {
	products := *s.seller.Products["Electronics"]

	if len(products) == 0 {
		fmt.Println("No products available")
		return
	}

	s.seller.Show()

	itemInput := extensionflow.UserInput("Enter product index: ")
	index, err := strconv.Atoi(itemInput)
	if err != nil {
		return
	}

	if index < 0 || index >= len(products) {
		fmt.Println("Invalid index")
		return
	}

	product := products[index]

	if s.seller.HasProduct(product.Name) && s.customer.HasEnoughMoney(product.Price) {
		s.seller.Sell(index)
		s.customer.Pay(product.Price)
		s.customer.TakeProduct(product)
	}
}
