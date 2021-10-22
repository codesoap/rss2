package rss2

import (
	"encoding/xml"
	"fmt"
)

// TextInput represents a Channel's textInput element. All sub-elements
// must be present.
type TextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

// NewTextInput creates a new TextInput element.
func NewTextInput(title, description, name, link string) (*TextInput, error) {
	if len(title) == 0 || len(description) == 0 || len(name) == 0 ||
		len(link) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewTextInput()`)
	}
	return &TextInput{
		XMLName:     xml.Name{Local: `textInput`},
		Title:       title,
		Description: description,
		Name:        name,
		Link:        link,
	}, nil
}
