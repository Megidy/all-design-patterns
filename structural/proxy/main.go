package main

import "log"

func main() {
	doc1 := NewProxyDocument(true)
	doc2 := NewProxyDocument(false)

	doc1.Display()
	doc2.Display()
}

type Proxy interface {
	Display()
}

type ProxyDocument struct {
	hasAccess bool
}

func NewProxyDocument(hasAccess bool) Proxy {
	return &ProxyDocument{
		hasAccess: hasAccess,
	}
}

func (pd *ProxyDocument) Display() {
	if pd.hasAccess {
		log.Printf("Displayed")
	} else {
		log.Printf("Access denied")
	}
}
