package main

import (
	"log"
	"pnode/internal/app"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Run()

}
