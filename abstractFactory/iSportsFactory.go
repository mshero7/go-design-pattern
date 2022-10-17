package abstractFactory

import "fmt"

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "nike" {
		return &Nike{}, nil
	} else if brand == "adidas" {
		return &Adidas{}, nil
	}

	return nil, fmt.Errorf("wrong Brand Type passed")
}
