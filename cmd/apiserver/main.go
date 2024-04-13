package main

import (
	"log"

	"github.com/abdirakhman/comp-prog-kz/internal/app/apiserver"
)

func main() {
	config := apiserver.GetConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
