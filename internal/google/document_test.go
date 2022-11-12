package google_test

import (
	"github.com/sebast26/txt2gdoc/internal/google"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDocument(t *testing.T) {
	t.Run("location", func(t *testing.T) {
		// when
		doc := google.NewDocument("some-google-document-id", "Title of the document")

		// then
		assert.Equal(t, "some-google-document-id", doc.ID)
		assert.Equal(t, "Title of the document", doc.Title)
		assert.Equal(t, "https://docs.google.com/document/d/some-google-document-id/edit", doc.Location)
	})
}
