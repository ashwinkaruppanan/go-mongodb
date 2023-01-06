package main

import (
	"fmt"

	"github.com/ashwin/go-mongodb/router"
)

func main() {
	fmt.Printf("go-mongodb API \n\n")
	fmt.Println("starting server...")

	router.Router()
}
