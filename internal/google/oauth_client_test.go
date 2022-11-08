package google_test

import (
	"github.com/sebast26/txt2gdoc/internal/google"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOAuthClient(t *testing.T) {
	t.Run("credential file not found", func(t *testing.T) {
		// when
		client, err := google.NewOAuthClient("non_existing_path")

		// then
		assert.ErrorContains(t, err, "reading client secret file failed")
		assert.Empty(t, client)
	})
}
