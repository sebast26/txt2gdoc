package main

import (
	"fmt"
	"github.com/sebast26/txt2gdoc/internal/stdin"
)

func main() {
	buf, err := stdin.ReadStdin()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read from stdin: %s", string(buf))
}
