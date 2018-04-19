package main

import (
	"fmt"
	"github.com/xrash/pgn2text/pkg/pgn2text"
	"os"
)

func main() {
	s, err := pgn2text.Do(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
