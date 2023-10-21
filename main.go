package main

import (
	"log"
)

// Create a basic URL shortner with map

func main() {
	err := execute()
	if err != nil {
		log.Fatal(err.Error())
	}

}
