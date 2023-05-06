package main

import (
	"log"
	"profile/scripts"
)

func main() {
	if err := scripts.Edit(); err != nil {
		log.Fatalf("%+v", err)
	}
}
