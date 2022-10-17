package abstractFactory

type IShirt interface {
	setShirtLogo(logo string)
	setShirtSize(size int)
	GetShirtLogo() string
	GetShirtSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) GetShirtLogo() string {
	return s.logo
}

func (s *Shirt) GetShirtSize() int {
	return s.size
}

func (s *Shirt) setShirtSize(size int) {
	s.size = size
}

func (s *Shirt) setShirtLogo(logo string) {
	s.logo = logo
}
