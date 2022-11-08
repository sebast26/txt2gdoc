package google

import (
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type OAuthClient struct {
	httpClient *http.Client
}

func NewOAuthClient(credentialsFilePath string) (*OAuthClient, error) {
	_, err := os.ReadFile(credentialsFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "reading client secret file failed")
	}

	return &OAuthClient{}, nil
}
