package builder

// concrete builder
type Normal struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *Normal {
	return &Normal{}
}

func (n *Normal) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *Normal) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *Normal) setNumFloor() {
	n.floor = 2
}

func (n *Normal) getHouse() House {
	return House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}

}
