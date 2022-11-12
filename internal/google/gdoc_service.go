package google

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
	"net/http"
	"strings"
	"time"
)

// DocumentService represents service that manipulates Google documents.
type DocumentService struct {
	gdoc *docs.Service
}

// NewDocumentService creates new DocumentService.
func NewDocumentService(oauthClient *http.Client) (*DocumentService, error) {
	ctx := context.Background()
	srv, err := docs.NewService(ctx, option.WithHTTPClient(oauthClient))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create GDocs client")
	}
	return &DocumentService{gdoc: srv}, nil
}

// CreateDocument creates new Google document with prefixed title and content.
func (s *DocumentService) CreateDocument(prefix, content string) (Document, error) {
	title := createDocumentTitle(prefix, content)
	doc, err := s.gdoc.Documents.Create(&docs.Document{Title: title}).Do()
	if err != nil {
		return Document{}, errors.Wrap(err, "unable to create a document")
	}

	md := docs.Dimension{Magnitude: 20.0, Unit: "PT"}
	style := docs.DocumentStyle{MarginLeft: &md, MarginRight: &md, MarginTop: &md}
	styleRequest := docs.UpdateDocumentStyleRequest{DocumentStyle: &style, Fields: "marginTop,marginLeft,marginRight"}
	ins1 := docs.InsertTextRequest{Text: content, Location: &docs.Location{Index: 1}}

	requests := []*docs.Request{
		{UpdateDocumentStyle: &styleRequest},
		{InsertText: &ins1},
	}
	_, err = s.gdoc.Documents.BatchUpdate(doc.DocumentId, &docs.BatchUpdateDocumentRequest{Requests: requests}).Do()
	if err != nil {
		return Document{}, errors.Wrap(err, "batch update failed")
	}
	return NewDocument(doc.DocumentId, doc.Title), nil
}

func createDocumentTitle(prefix, content string) string {
	i := strings.Index(content, "\n")
	if i == -1 || i > maxContentLengthToTitle {
		i = maxContentLengthToTitle
	}

	currentTime := time.Now()
	if prefix == "" {
		return fmt.Sprintf("%s - %s...", currentTime.Format("2006-01-02"), content[:i])
	}
	return fmt.Sprintf("%s %s - %s...", prefix, currentTime.Format("2006-01-02"), content[:i])
}

const (
	// maxContentLengthToTitle specifies maximum number of characters that could be included inside document title
	maxContentLengthToTitle = 30
)
