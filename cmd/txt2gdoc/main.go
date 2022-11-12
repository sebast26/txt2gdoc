package main

import (
	"github.com/sebast26/txt2gdoc/internal/google"
	"github.com/sebast26/txt2gdoc/internal/stdin"
)

func main() {
	buf, err := stdin.ReadStdin()
	if err != nil {
		panic(err)
	}

	client, err := google.NewOAuthClient("credentials.json", "token.json")
	if err != nil {
		panic(err)
	}
	service, err := google.NewDocumentService(client.HttpClient)
	if err != nil {
		panic(err)
	}
	doc, err := service.CreateDocument("", string(buf))
	if err != nil {
		panic(err)
	}

	println(doc.Location)
}
