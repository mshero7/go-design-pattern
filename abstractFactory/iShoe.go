package abstractFactory

type IShoe interface {
	setShoeLogo(logo string)
	setShoeSize(size int)
	GetShoeLogo() string
	GetShoeSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setShoeLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setShoeSize(size int) {
	s.size = size
}

func (s *Shoe) GetShoeLogo() string {
	return s.logo
}

func (s *Shoe) GetShoeSize() int {
	return s.size
}
