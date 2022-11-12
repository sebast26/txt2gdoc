package google

import "fmt"

// Document represents Google document.
type Document struct {
	ID       string
	Title    string
	Location string
}

// NewDocument creates new Google Document.
func NewDocument(ID, Title string) Document {
	return Document{
		ID:       ID,
		Title:    Title,
		Location: fmt.Sprintf(documentLocationTemplate, ID),
	}
}

const (
	// documentLocationTemplate is a template used by Google Docs to access document by ID from the browser
	documentLocationTemplate = "https://docs.google.com/document/d/%s/edit"
)
