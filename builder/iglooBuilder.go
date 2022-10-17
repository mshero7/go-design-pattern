package builder

// concrete builder
type Igloo struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *Igloo {
	return &Igloo{}
}

func (i *Igloo) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *Igloo) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *Igloo) setNumFloor() {
	i.floor = 1
}

func (i *Igloo) getHouse() House {
	return House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}

}
