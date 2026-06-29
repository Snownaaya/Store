package domain

import (
	"fmt"
	"strconv"

	"store/extensionflow"
)

type Shop struct {
	customer    *Customer
	seller      *Seller
	commands    map[string]func()
	isWork      bool
	commandMenu CommandMenu
}

type Command struct {
	InputNeeded string
	Description string
	Action      func()
}

type CommandMenu struct {
	commands []Command
}

func NewShop() *Shop {
	shop := &Shop{
		customer:    newCustomer(),
		seller:      NewSeller(),
		commandMenu: CommandMenu{},
		isWork:      true,
	}

	menu := CommandMenu{}

	menu.add(Command{
		InputNeeded: "1",
		Description: "ShowInfo",
		Action:      func() { shop.seller.Show() },
	})

	menu.add(Command{
		InputNeeded: "2",
		Description: "Sell",
		Action:      func() { shop.sellProductFlow() },
	})

	menu.add(Command{
		InputNeeded: "3",
		Description: "Inventory",
		Action:      func() { shop.customer.Show() },
	})

	menu.add(Command{
		InputNeeded: "4",
		Description: "Exit",
		Action:      func() { shop.isWork = false },
	})

	shop.commandMenu = menu
	shop.commands = menu.toActionMap()

	return shop
}

func (s *Shop) Run() {
	for s.isWork {
		s.commandMenu.printMenu()
		input := extensionflow.UserInput("Enter command: ")

		if cmd, ok := s.commands[input]; ok {
			cmd()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func (m *CommandMenu) toActionMap() map[string]func() {
	actionMap := make(map[string]func())
	for _, cmd := range m.commands {
		actionMap[cmd.InputNeeded] = cmd.Action
	}

	return actionMap
}

func (c *CommandMenu) printMenu() {
	for _, cmd := range c.commands {
		fmt.Println(cmd.InputNeeded, cmd.Description)
	}
}

func (c *CommandMenu) add(command Command) {
	c.commands = append(c.commands, command)
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
