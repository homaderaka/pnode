package main

import (
	"context"
	"log"
	"pnode/internal/app"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	c := context.Background()

	a.Run(c)

}
