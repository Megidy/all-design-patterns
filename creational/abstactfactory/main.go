package main

import (
	"fmt"
	"log"
)

func main() {
	factory, err := NewGuiFactory("mac")
	if err != nil {
		log.Fatal(err)
	}
	RenderUI(factory)
}

const (
	windows = "windows"
	mac     = "mac"
)

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsButton struct{}

func (b *WindowsButton) Paint() {
	log.Println("windows style button")
}

type WindowsCheckbox struct{}

func (c *WindowsCheckbox) Paint() {
	log.Println("windows style checkbox")
}

type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type MacButton struct{}

func (b *MacButton) Paint() {
	log.Println("mac style button")
}

type MacCheckbox struct{}

func (c *MacCheckbox) Paint() {
	log.Println("mac style checkbox")
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

func NewGuiFactory(platform string) (GUIFactory, error) {
	switch platform {
	case windows:
		return &WindowsFactory{}, nil
	case mac:
		return &MacFactory{}, nil
	default:
		return nil, fmt.Errorf("invalid platform")
	}
}

func RenderUI(factory GUIFactory) {
	btn := factory.CreateButton()
	box := factory.CreateCheckbox()

	btn.Paint()
	box.Paint()

}
