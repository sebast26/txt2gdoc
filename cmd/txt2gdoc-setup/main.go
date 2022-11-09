package main

import (
	"fmt"
	"github.com/sebast26/txt2gdoc/internal/google"
)

func main() {
	_, err := google.NewOAuthClient("credentials.json", "token.json")
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
	fmt.Println("txt2gdoc setup completed")

}
