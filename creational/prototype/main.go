package main

import "fmt"

func main() {

	original := NewTextDocument()
	clone := original.Clone().(*TextDocument)

	clone.Metadata["hello"] = "clone"

	fmt.Println("Original:", original.Metadata["hello"])
	fmt.Println("Clone:   ", clone.Metadata["hello"])
}

type TextDocument struct {
	Title    string
	Content  string
	Metadata map[string]string
}

func NewTextDocument() *TextDocument {
	return &TextDocument{
		Title:   "title",
		Content: "content",
		Metadata: map[string]string{
			"hello": "world",
		},
	}
}

func (d *TextDocument) DoStuff() {}

func (d *TextDocument) Clone() Document {
	var resp = &TextDocument{
		Title:    d.Title,
		Content:  d.Content,
		Metadata: make(map[string]string),
	}

	for k, v := range d.Metadata {
		resp.Metadata[k] = v
	}

	return resp
}

type Document interface {
	DoStuff()
	Clone() Document
}
