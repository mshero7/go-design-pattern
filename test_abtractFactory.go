package main

import (
	"fmt"
	abs "go-example/design-pattern/abstractFactory"
)

func TestAbstractFactory() {
	adidasFactory, _ := abs.GetSportsFactory("adidas")
	nikeFactory, _ := abs.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s abs.IShoe) {
	fmt.Printf("Logo: %s", s.GetShoeLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetShoeSize())
	fmt.Println()
}

func printShirtDetails(s abs.IShirt) {
	fmt.Printf("Logo: %s", s.GetShirtLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetShirtSize())
	fmt.Println()
}
