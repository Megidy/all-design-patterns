package main

import (
	"log"
)

type Config struct {
	Field1 string
	Field2 string
}

// add mutex or sync once to make this concurrency safe
var globalConfig *Config

func main() {
	c1 := NewConfig()
	c2 := NewConfig()

	log.Printf("c1: %p", c1)
	log.Printf("c2: %p", c2)
}

func NewConfig() *Config {
	if globalConfig != nil {
		return globalConfig
	} else {
		globalConfig = &Config{
			Field1: "field1",
			Field2: "field2",
		}
		return globalConfig
	}
}
