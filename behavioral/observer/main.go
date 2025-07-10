package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	controller := NewController()
	service := NewService()

	controller.AddListener(service)

	go func() {
		<-time.After(time.Second)
		controller.RemoveListener(service)
	}()

	for range 10 {
		newEvent := &Event{
			data: "New data",
		}

		controller.Notify(newEvent)

		time.Sleep(time.Millisecond * 400)
	}

}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DoThing() {}

func (s *Service) NotifyCallBack(event *Event) {
	log.Printf("service received and hanled event with data: %v", event.data)
}

type Controller struct {
	mu        sync.Mutex
	observers map[Observer]struct{}
}

func NewController() *Controller {
	return &Controller{
		mu:        sync.Mutex{},
		observers: make(map[Observer]struct{}),
	}
}

func (c *Controller) AddListener(obs Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.observers[obs] = struct{}{}
	log.Printf("added observer")
}
func (c *Controller) RemoveListener(obs Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.observers, obs)
	log.Printf("removed observer")
}
func (c *Controller) Notify(event *Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for observer := range c.observers {
		observer.NotifyCallBack(event)
	}
}

type Event struct {
	data any
}

type Observer interface {
	NotifyCallBack(*Event)
}

type Subject interface {
	AddListener(Observer)
	RemoveListener(Observer)
	Notify(*Event)
}
