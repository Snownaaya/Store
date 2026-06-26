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
}

const (
	ShowPotionCommand string = "1"
	SellCommand       string = "2"
	BuyerCommand      string = "3"
	ExitCommand       string = "4"
)

func NewShop() *Shop {
	shop := &Shop{
		customer: newCustomer(),
		seller:   NewSeller(),
	}

	return shop
}

func (s *Shop) Run() {
	isWork := true

	for isWork {
		fmt.Println("1 - ShowPotions")
		fmt.Println("2 - Sell")
		fmt.Println("3 - Inventory")
		fmt.Println("4 - Exit")

		input := extensionflow.UserInput("Enter command: ")

		s.commands = map[string]func(){
			ShowPotionCommand: func() {
				s.seller.Show()

			},
			SellCommand: func() {
				s.seller.Show()

				itemInput := extensionflow.UserInput("Enter product index: ")
				index, _ := strconv.Atoi(itemInput)

				products := *s.seller.Products["Electronics"]
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
			},
			BuyerCommand: func() {
				s.customer.Show()
			},
			ExitCommand: func() {
				isWork = false
			},
		}

		if cmd, ok := s.commands[input]; ok {
			cmd()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
